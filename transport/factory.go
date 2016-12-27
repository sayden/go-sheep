package transport

import "github.com/sayden/go-sheep"

type Type string

const (
	GRPC Type = "GRPC"
	REST Type = "REST"
)

func NewClient(t Type) go_sheep.Transporter {
	switch t {
	case GRPC:
		return &grpcClient{}
	case REST:
		panic("Not implemented yet")
	default:
		return &grpcClient{}
	}
}

func NewServer(t Type) go_sheep.SWIMClient{
	switch t {
	case GRPC:
		return &grpcServer{}
	case REST:
		panic("Not implemented yet")
	}
}

func New(t Type) go_sheep.Transporter {

}