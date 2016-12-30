package transport

import "github.com/sayden/go-sheep"

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

func New(t Type) Transporter {
	switch t {
	case GRPC:
		return &GRPCTransport{}
	case REST:
		panic("Not implemented yet")
	default:
		panic("Not implemented yet")
	}
}
