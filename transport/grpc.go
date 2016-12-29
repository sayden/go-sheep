package transport

import "github.com/sayden/go-sheep"

type GRPCTransport struct {
	GRPCServer
	GRPCClient
	go_sheep.SWIM
}
