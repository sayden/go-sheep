package transport

import "github.com/sayden/go-sheep"

type grpcTransport struct {

}

func (g *grpcTransport) Ping(s *go_sheep.State, a string) (*go_sheep.State, error) {
	panic("not implemented")
}

