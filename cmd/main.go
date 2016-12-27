package main

import (
	"os"

	"github.com/sayden/go-sheep"
	"github.com/uber-go/zap"
	"github.com/urfave/cli"
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
			Name:    "join",
			Aliases: []string{"j"},
			Usage:   "Join a cluster",
			Action:  joinAction,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:   "port, p",
					EnvVar: "GO_SHEEP_PEER_PORT",
					Value:  8080,
					Usage:  "Port of peer to join",
				},
				cli.StringFlag{
					Name:   "bind, b",
					Usage:  "IP to bind to",
					Value:  "0.0.0.0",
					EnvVar: "GO_SHEEP_PEER_IP",
				},
			},
		},
		{
			Name:    "agent",
			Aliases: []string{"a"},
			Usage:   "Launch agent",
			Action:  agentAction,
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
	}

	app.Run(os.Args)
}

func getNewUUID() string {
	return ""
}
