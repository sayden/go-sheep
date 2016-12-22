package transport

import (
	"context"

	"sync"

	"github.com/sayden/go-sheep"
	"google.golang.org/grpc"
)

type grpcTransport struct{
	go_sheep.Transporter
}

func (g *grpcTransport) Ping(s *go_sheep.State, a string) (state *go_sheep.State, err error) {
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(a, grpc.WithInsecure())
	if err != nil {
		return
	}

	client := go_sheep.NewSWIMClient(conn)
	state, err = client.Ping(context.Background(), s)

	return
}

func (g *grpcTransport) IndirectPing(s *go_sheep.State, delegatedNodes []string, t string) (states []*go_sheep.State, errors error) {
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

	sheepError := &go_sheep.Error{
		Errs:    make([]error, len(errCh)),
		File:    "grpc.go",
		Method:  "IndirectPing",
		Package: "transport",
	}

	var i int
	for err := range errCh {
		sheepError.Errs[i] = err
		i++
	}

	errors = sheepError

	return
}
