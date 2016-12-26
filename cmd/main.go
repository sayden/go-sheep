package main

import (
	"fmt"
	"os"
	"time"

	"github.com/concourse/concourse/src/github.com/golang/protobuf/ptypes"
	"github.com/micro/cli"
	"github.com/sayden/go-sheep"
	"github.com/sayden/go-sheep/failure_detector"
	"github.com/uber-go/zap"
)

var logger zap.Logger = zap.New(zap.NewTextEncoder())

var state *go_sheep.State

var loopTime int = 5

var currentNode *go_sheep.Node

func main() {

	state = &go_sheep.State{
		Nodes: make([]*go_sheep.Node, 0),
	}

	state.Nodes = []*go_sheep.Node{currentNode}

	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"

	app.Commands = []cli.Command{
		{
			Name:    "join",
			Aliases: []string{"j"},
			Usage:   "Join a cluster",
			Action: func(c *cli.Context) {
				fmt.Println("added task: ", c.Args().First())
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:   "port, p",
					EnvVar: "GO_SHEEP_AGENT_PORT",
					Value:  8080,
					Usage:  "Port to listen for other nodes",
				},
				cli.StringFlag{
					Name:   "bind, b",
					Usage:  "IP to bind to",
					Value:  "0.0.0.0",
					EnvVar: "GO_SHEEP_BIND_IP",
				},
			},
		},
		{
			Name:    "agent",
			Aliases: []string{"a"},
			Usage:   "Launch agent",
			Action: func(c *cli.Context) {
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

				//init listener
				//TODO

				loop()
			},
		},
	}

	app.Run(os.Args)
}

func loop() {
	swim := failure_detector.NewSwim()

	for {
		// Get a randomized node from the state
		target, err := swim.GetRandomizedTarget(state, currentNode)
		if err != nil {
			logger.Error("Could not get randomized target", zap.Error(err), zap.Object("State", state))
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

func getNewUUID() string {
	return ""
}
