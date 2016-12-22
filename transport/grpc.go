package transport

import (
	"context"

	"github.com/sayden/go-sheep"
	"google.golang.org/grpc"
)

type grpcTransport struct {
}

func (g *grpcTransport) Ping(s *go_sheep.State, a string) (state *go_sheep.State, err error) {
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(a, grpc.WithInsecure())
	if err != nil {
		return
	}

	client := go_sheep.NewPingerClient(conn)
	state, err = client.Ping(context.Background(), s)

	return
}

func (g *grpcTransport) IndirectPing(s *go_sheep.State, a []string, t string) ([]*go_sheep.State, error) {
	in := make(chan struct{
		States []*go_sheep.State
		Error error
	}, len(a))

	for _, node := range a {
		go func(addr string) {
			//Ask to ping node t
		}(node)
	}
}
