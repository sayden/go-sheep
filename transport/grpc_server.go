package transport

import (
	"net"

	"fmt"

	"github.com/grafana/grafana/pkg/cmd/grafana-cli/logger"
	"github.com/sayden/go-sheep"
	"github.com/sayden/go-sheep/failure_detector"
	"github.com/uber-go/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	Server
	logger zap.Logger
}

func (g *GRPCServer) ServeJoin(n *go_sheep.Node) (*go_sheep.State, error) {
	go_sheep.CurrentState.Lock()
	defer go_sheep.CurrentState.Unlock()

	logger.Info("Join request arrived of\n", zap.Object("node", n))

	newState, err := failure_detector.Swim{}.MergeState(&go_sheep.State{
		Nodes: []*go_sheep.Node{n},
	}, go_sheep.CurrentState.State)
	if err != nil {
		return nil, err
	}

	go_sheep.CurrentState.State = newState
	return go_sheep.CurrentState.State, nil
}

func (g *GRPCServer) ServePing(s *go_sheep.State) (*go_sheep.State, error) {
	go_sheep.CurrentState.Lock()
	defer go_sheep.CurrentState.Unlock()


	newState, err := failure_detector.Swim{}.MergeState(s, go_sheep.CurrentState.State)
	if err != nil {
		return nil, err
	}

	go_sheep.CurrentState.State = newState
	return go_sheep.CurrentState.State, nil
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

func (g *GRPCServer) Serve(address string, quit chan struct{}, logger zap.Logger) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	go_sheep.RegisterSWIMServer(s, &GRPCServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		select {
		case <-quit:
			logger.Debug("Stopping GRPC server")
			s.Stop()
		}
	}()

	go func(logger zap.Logger) {
		if err := s.Serve(lis); err != nil {
			logger.Error("failed to serve:", zap.Error(err))
		}
	}(logger)

	return nil
}
