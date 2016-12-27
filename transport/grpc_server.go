package transport

import (
	"context"

	"net"

	"fmt"

	"errors"

	"github.com/sayden/go-sheep"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	go_sheep.SWIMServer
}

func NewServer(p string) error {
	lis, err := net.Listen("tcp", p)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	go_sheep.RegisterSWIMServer(s, &grpcServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
}

func (swim *grpcServer) Ping(ctx context.Context, state *go_sheep.State) (*go_sheep.State, error) {
	return nil, errors.New("Not implemented yet")
}

func (swim *grpcServer) DelegateCheck(ctx context.Context, state *go_sheep.DelegateCheckRequest) (*go_sheep.State, error) {
	return nil, errors.New("Not implemented yet")
}

func (swim *grpcServer) Join(_ context.Context, state *go_sheep.Node) (*go_sheep.State, error) {
	fmt.Printf("%#v\n", state)
	return nil, nil
}
