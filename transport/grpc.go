package transport

import (
	"github.com/sayden/go-sheep"
	"github.com/uber-go/zap"
)

func NewGRPCTransport(l zap.Logger) *GRPCTransport {
	return &GRPCTransport{
		GRPCServer:GRPCServer{
			logger:l,
		},
		GRPCClient:GRPCClient{
			logger:l,
		},
	}
}

type GRPCTransport struct {
	GRPCServer
	GRPCClient
	go_sheep.SWIM
}
