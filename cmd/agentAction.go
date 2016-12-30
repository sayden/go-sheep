package main

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
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
		Uuid:     getNewUUID(),
		LastSeen: now,
	}

	//retrieve transport type
	transport := transport.New(transport.Type(c.String("transport")))

	//retrieve failure detector
	failureDetector := failure_detector.NewFailureDetector()

	//Launch state handler
	reqCh, quit := go_sheep.LaunchStateHandler(currentNode, failureDetector, logger)

	//init server
	err = transport.Serve(c.String("port"))
	if err != nil {
		logger.Fatal("Could not start server", zap.Error(err))
	}

	//launch loop
	loop(c)
}

func loop(c *cli.Context) {
	swim := failure_detector.NewSwim(
		transport.New(transport.Type(c.String("transport"))),
	)

	for {
		// Get a randomized node from the state
		target, err := swim.GetRandomizedTarget(state, currentNode)
		if err != nil {
			logger.Error("Could not get randomized target",
				zap.Error(err),
				zap.Object("State", state),
			)

			goto end
		}

		// Ping it
		err = swim.CheckNode(state, target.Address, currentNode.Address)
		switch errType := err.(type) {
		case go_sheep.NetworkError:
			logger.Warn("Could not contact node", zap.Error(err))

			checkers, err := swim.GetCheckers(state, target, currentNode, 2)
			if err != nil {
				logger.Error("Error getting node checkers",
					zap.Error(err),
					zap.Int("Checkers", 2),
					zap.Object("TargetNode", target),
					zap.Object("CurrentNode", currentNode),
					zap.Object("State", state))
				goto end
			}

			receivedStates, err := swim.IndirectPing(state, checkers, target)
			if err != nil {
				logger.Error("Error getting node checkers",
					zap.Error(err),
					zap.Int("CheckersNumber", 2),
					zap.Object("Checkers", checkers),
					zap.Object("TargetNode", target),
					zap.Object("CurrentNode", currentNode),
					zap.Object("State", state))
			}

			for _, newState := range receivedStates {
				res, err := swim.MergeState(state, newState)
				if err != nil {
					logger.Error("Error merging states",
						zap.Error(err),
						zap.Object("State", state))
					zap.Object("ReceivedState", newState)
					goto end
				}

				state = res
			}
		case error:
			logger.Error("Error getting node to check", zap.Error(errType))
		default:
			logger.Error("Error getting node to check", zap.Error(errType))
		}

	end:
		time.Sleep(time.Second * time.Duration(loopTime))
	}
}
