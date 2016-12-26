package transport

import (
	"context"

	"sync"

	"github.com/sayden/go-sheep"
	"google.golang.org/grpc"
)

type grpcClient struct {
	go_sheep.Transporter
}

func NewGRPCTransport() go_sheep.Transporter {
	return &grpcClient{}
}

func (g *grpcClient) Ping(s *go_sheep.State, a string) (state *go_sheep.State, err error) {
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(a, grpc.WithInsecure())
	if err != nil {
		return
	}

	client := go_sheep.NewSWIMClient(conn)
	state, err = client.Ping(context.Background(), s)

	return
}

func (g *grpcClient) IndirectPing(s *go_sheep.State, delegatedNodes []string, t string) (states []*go_sheep.State, err error) {
	out := make(chan *go_sheep.State, len(delegatedNodes))
	errCh := make(chan error, len(delegatedNodes))

	var wg sync.WaitGroup
	wg.Add(len(delegatedNodes))

	for _, node := range delegatedNodes {
		go func(out chan<- *go_sheep.State, s *go_sheep.State, delegatedNode, t string) {
			defer wg.Done()

			var conn *grpc.ClientConn
			conn, err := grpc.Dial(delegatedNode, grpc.WithInsecure())
			if err != nil {
				errCh <- go_sheep.NewCheckError(
					go_sheep.Node{Address: delegatedNode},
					go_sheep.Node{Address: t},
					err,
				)
				return
			}

			client := go_sheep.NewSWIMClient(conn)

			var res *go_sheep.State
			res, err = client.DelegateCheck(context.Background(), &go_sheep.DelegateCheckRequest{
				State:  s,
				Target: t,
			})

			if err != nil {
				errCh <- go_sheep.NewCheckError(
					go_sheep.Node{Address: delegatedNode},
					go_sheep.Node{Address: t},
					err,
				)
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

func (swim *grpcClient) Join(in, targetServer *go_sheep.Node) (state *go_sheep.State, err error) {
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(targetServer.Address, grpc.WithInsecure())
	if err != nil {
		return
	}

	client := go_sheep.NewSWIMClient(conn)
	state, err = client.Join(context.Background(), in)

	return
}

type grpcServer struct {
	go_sheep.SWIMServer
}

func (swim *grpcServer) Ping(context.Context, *go_sheep.State) (*go_sheep.State, error) {

}

func (swim *grpcServer) DelegateCheck(context.Context, *go_sheep.DelegateCheckRequest) (*go_sheep.State, error) {

}

func (swim *grpcServer) Join(context.Context, *go_sheep.Node) (*go_sheep.State, error) {

}
