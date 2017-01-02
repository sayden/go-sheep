package transport

import (
	"github.com/sayden/go-sheep"
	"github.com/uber-go/zap"
)

type Transporter interface {
	Client
	Server
	go_sheep.SWIM
}

type Type string

const (
	GRPC Type = "GRPC"
	REST Type = "REST"
)

func New(t Type, l zap.Logger) Transporter {
	switch t {
	case GRPC:
		return NewGRPCTransport(l)
	case REST:
		panic("Not implemented yet")
	default:
		panic("Not implemented yet")
	}
}
