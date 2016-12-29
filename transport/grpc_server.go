package transport

import (
	"net"

	"fmt"

	"github.com/sayden/go-sheep"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct{}

func (g *GRPCServer) ServeJoin(*go_sheep.Node) (*go_sheep.State, error) {
	panic("implement me")
}

func (g *GRPCServer) ServePing(s *go_sheep.State) (*go_sheep.State, error) {
	newState := g.swim.MergeState()
}

func (g *GRPCServer) ServeIndirectCheck(*go_sheep.DelegateCheckRequest) (*go_sheep.State, error) {
	panic("implement me")
}

func (g *GRPCServer) DelegateCheck(_ context.Context, req *go_sheep.DelegateCheckRequest) (*go_sheep.State, error) {
	return g.ServeIndirectCheck(req)
}

func (g *GRPCServer) Join(_ context.Context, n *go_sheep.Node) (*go_sheep.State, error) {
	return g.ServeJoin(n)
}

func (g *GRPCServer) Ping(_ context.Context, s *go_sheep.State) (*go_sheep.State, error) {
	return g.ServePing(s)
}

func (g *GRPCServer) Serve(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	go_sheep.RegisterSWIMServer(s, &GRPCServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
