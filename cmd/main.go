package main

import (
	"os"

	"github.com/sayden/go-sheep"
	"github.com/uber-go/zap"
	"github.com/urfave/cli"
)

var logger zap.Logger = zap.New(zap.NewTextEncoder())

var loopTime int = 5

var currentNode *go_sheep.Node

func main() {
	go_sheep.CurrentState = &go_sheep.SafeState{
		State: &go_sheep.State{
			Nodes: make([]*go_sheep.Node, 1),
		},
	}

	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "transport, t",
			EnvVar: "GO_SHEEP_TRANSPORT",
			Value:  "GRPC",
			Usage:  "Choose transport for communication. Only GRPC supported. REST will come eventually",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "agent",
			Aliases: []string{"a"},
			Usage:   "Launch agent",
			Action:  agentAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "port, p",
					EnvVar: "GO_SHEEP_AGENT_PORT",
					Value:  "8080",
					Usage:  "Port to listen for other nodes",
				},
				cli.StringFlag{
					Name:   "bind, b",
					Usage:  "IP to bind to",
					Value:  "0.0.0.0",
					EnvVar: "GO_SHEEP_BIND_IP",
				},
				cli.StringFlag{
					Name:   "target-server",
					EnvVar: "GO_SHEEP_TARGET_SERVER",
					Value:  "0.0.0.0",
					Usage:  "Server to join to",
				},
				cli.StringFlag{
					Name:   "target-port",
					EnvVar: "GO_SHEEP_TARGET_PORT",
					Usage:  "Port where the Server to join to is listening",
				},
			},
		},
	}

	app.Run(os.Args)
}

func getNewUUID() string {
	return ""
}
