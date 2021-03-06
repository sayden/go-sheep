package transport

import (
	"context"

	"google.golang.org/grpc"

	"sync"

	"github.com/sayden/go-sheep"
	"github.com/uber-go/zap"
)

type GRPCClient struct {
	Client
	logger zap.Logger
}

func (g *GRPCClient) DoIndirectPing(s *go_sheep.State, delegatedNodes []*go_sheep.Node, t *go_sheep.Node) (states []*go_sheep.State, err error) {
	out := make(chan *go_sheep.State, len(delegatedNodes))
	errCh := make(chan error, len(delegatedNodes))

	var wg sync.WaitGroup
	wg.Add(len(delegatedNodes))

	for _, node := range delegatedNodes {
		go func(out chan<- *go_sheep.State, s *go_sheep.State, delegatedNode, t *go_sheep.Node) {
			defer wg.Done()

			var conn *grpc.ClientConn
			conn, err := grpc.Dial(delegatedNode.Address, grpc.WithInsecure())
			if err != nil {
				errCh <- go_sheep.NewCheckError(delegatedNode,t,err)
				return
			}

			client := go_sheep.NewSWIMClient(conn)

			var res *go_sheep.State
			res, err = client.DelegateCheck(context.Background(), &go_sheep.DelegateCheckRequest{
				State:  s,
				Target: t,
			})

			if err != nil {
				errCh <- go_sheep.NewCheckError(delegatedNode,t,err)
				return
			}

			out <- res
		}(out, s, node, t)
	}

	wg.Wait()

	close(out)
	close(errCh)

	states = make([]*go_sheep.State, 0)
	for state := range out {
		states = append(states, state)
	}

	errors := go_sheep.NewErrors(len(errCh))

	var i int
	for err := range errCh {
		(*errors)[i] = err
		i++
	}

	err = errors

	return
}

func (g *GRPCClient) DoJoin(in, targetServer *go_sheep.Node) (state *go_sheep.State, err error) {
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(targetServer.Address, grpc.WithInsecure())
	if err != nil {
		return
	}

	client := go_sheep.NewSWIMClient(conn)
	state, err = client.Join(context.Background(), in)

	return
}

func (g *GRPCClient) DoPing(s *go_sheep.State, a string) (state *go_sheep.State, err error) {
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(a, grpc.WithInsecure())
	if err != nil {
		return
	}

	client := go_sheep.NewSWIMClient(conn)
	state, err = client.Ping(context.Background(), s)

	return
}
