package go_sheep

import (
	"sync"
	"github.com/uber-go/zap"
)

type SafeState struct {
	*State
	sync.RWMutex
}

var CurrentState *SafeState

type StateAction byte

const (
	MERGE StateAction = iota
)

type Request struct {
	Action     StateAction
	ResponseCh chan *Response
	Payload    interface{}
}

type Response struct {
	State *State
	Err   error
}

func LaunchStateHandler(n *Node, failureDetector SWIM, logger zap.Logger) (chan<- Request, chan<- struct{}) {
	in := make(chan Request, 0)
	quit := make(chan struct{})

	state := &State{
		Nodes: []*Node{n},
	}

	go func() {
		var err error

		for {
			select {
			case req := <-in:
				switch req.Action {
				case MERGE:
					newState := req.Payload.(*State)

					response := &Response{}

					state, err = failureDetector.MergeState(state, newState)
					if err != nil {
						logger.Error("Failed when trying to merge two states",
							zap.Error(err),
							zap.Object("new-state", newState),
							zap.Object("old-state", state),
						)

						response.Err = err
					}

					response.State = state

					req.ResponseCh <- response
				}
			case <-quit:
				break
			}
		}
	}()

	return in, quit
}
