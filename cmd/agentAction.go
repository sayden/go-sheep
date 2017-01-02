package main

import (
	"fmt"
	"time"

	"os"
	"os/signal"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sayden/go-sheep"
	"github.com/sayden/go-sheep/failure_detector"
	"github.com/sayden/go-sheep/transport"
	"github.com/uber-go/zap"
	"github.com/urfave/cli"
)

func agentAction(c *cli.Context) {
	//Configure current node
	now, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		panic(err)
	}

	currentNode = &go_sheep.Node{
		Address:  fmt.Sprintf("%s:%d", c.String("bind"), c.Int("port")),
		Uuid:     uuid.New().String(),
		LastSeen: now,
	}

	go_sheep.CurrentState.Lock()
	go_sheep.CurrentState.State.Nodes[0] = currentNode
	go_sheep.CurrentState.Unlock()

	//retrieve transport type
	transport := transport.New(transport.Type(c.GlobalString("transport")), logger)

	//retrieve failure detector
	failureDetector := failure_detector.NewFailureDetector()

	//Join to a target server
	if c.String("target-port") != "" {
		target := &go_sheep.Node{
			Address: fmt.Sprintf("%s:%s", c.String("target-server"), c.String("target-port")),
		}

		state, err := transport.DoJoin(currentNode, target)
		if err != nil {
			logger.Error("Could not join server", zap.Error(err))
		}

		go_sheep.CurrentState.Lock()
		go_sheep.CurrentState.State = state
		go_sheep.CurrentState.Unlock()
	}

	//Capture os signals to quit
	quit := make(chan struct{}, 0)

	//init server
	err = transport.Serve(fmt.Sprintf("%s:%s", c.String("bind"), c.String("port")), quit, logger)
	if err != nil {
		logger.Fatal("Could not start server", zap.Error(err))
	}

	//launch loop
	loop(c, failureDetector, transport, quit)
}

func loop(c *cli.Context, swim go_sheep.SWIM, transport transport.Transporter, quit chan struct{}) {
	quitOS := make(chan os.Signal, 1)
	signal.Notify(quitOS, os.Interrupt)

	for {
		select {
		case <-quitOS:
			logger.Info("Bye")
			close(quit)
			close(quitOS)
			return
		case <-time.After(time.Millisecond):
		}

		//logger.Debug("Sleeping")
		time.Sleep(time.Second * time.Duration(loopTime))

		// Get a randomized node from the state
		//logger.Info("Getting targets")

		go_sheep.CurrentState.RLock()
		tempState := go_sheep.CurrentState.State
		go_sheep.CurrentState.RUnlock()

		target, err := swim.GetRandomizedTarget(tempState, currentNode)
		if err != nil {
			logger.Error("Could not get randomized target",
				zap.Error(err),
				zap.Object("State", tempState),
			)

			continue
		}

		// Ping it
		var newState *go_sheep.State

		//logger.Info("Reading state")
		go_sheep.CurrentState.RLock()
		tempState = go_sheep.CurrentState.State
		go_sheep.CurrentState.RUnlock()

		//logger.Info("Doing ping")
		newState, err = transport.DoPing(tempState, target.Address)
		//logger.Info("Ping finished")

		if err != nil {
			logger.Warn("Could not contact node", zap.Error(err))

			checkers, err := swim.GetCheckers(tempState, target, currentNode, 2)
			if err != nil {
				logger.Error("Error getting node checkers",
					zap.Error(err),
					zap.Int("Checkers", 2),
					zap.Object("TargetNode", target),
					zap.Object("CurrentNode", currentNode))
				continue
			}

			receivedStates, err := transport.DoIndirectPing(tempState, checkers, target)
			if err != nil {
				logger.Error("Error getting node checkers",
					zap.Error(err),
					zap.Int("CheckersNumber", 2),
					zap.Object("Checkers", checkers),
					zap.Object("TargetNode", target),
					zap.Object("CurrentNode", currentNode))
				continue
			}

			for _, newState := range receivedStates {
				res, err := swim.MergeState(tempState, newState)
				if err != nil {
					logger.Error("Error merging states",
						zap.Error(err),
						zap.Object("State", tempState))
					continue
				}

				go_sheep.CurrentState.Lock()
				go_sheep.CurrentState.State = res
				go_sheep.CurrentState.Unlock()
				continue
			}
		}

		//No error?
		//logger.Info("No error")

		logger.Info("Current state", zap.Object("state", newState))

		go_sheep.CurrentState.Lock()
		go_sheep.CurrentState.State = newState
		go_sheep.CurrentState.Unlock()
	}
}
