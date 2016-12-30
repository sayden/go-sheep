package failure_detector

import (
	"errors"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/sayden/go-sheep"
	"github.com/sayden/go-sheep/transport"
)

type swim struct {}

func (swim *swim) GetRandomizedTarget(s *go_sheep.State, currentNodeInfo *go_sheep.Node) (n *go_sheep.Node, err error) {
	if len(s.Nodes) < 2 {
		return nil, errors.New("Only one node to randomize")
	}

	rand.Seed(time.Now().UnixNano())

	for {
		n = s.Nodes[rand.Intn(len(s.Nodes))]
		if n != currentNodeInfo {
			return
		}
	}
}

func (swim *swim) GetCheckers(s *go_sheep.State, targetNode *go_sheep.Node, currentNode *go_sheep.Node, n int) (foundNodes []*go_sheep.Node, err error) {
	if len(s.Nodes) < n+1 {
		return nil, errors.New("Not enough nodes in state")
	}

	rand.Seed(time.Now().UnixNano())

	foundNodes = make([]*go_sheep.Node, n)
	var arPos int
	for {
		candidate := s.Nodes[rand.Intn(len(s.Nodes))]
		if candidate != currentNode && candidate != targetNode {

			var found bool
			for _, foundNode := range foundNodes {
				if foundNode != nil && foundNode == candidate {
					found = true
					break
				}
			}

			if !found {
				foundNodes[arPos] = candidate
				arPos++
			}
		}

		if arPos >= n-1 {
			return
		}
	}
}

func (swim *swim) MergeState(as, bs *go_sheep.State) (newState *go_sheep.State, err error) {
	newState = &go_sheep.State{
		Nodes: make([]*go_sheep.Node, 0),
	}

	notFound := true
	var aTime, bTime time.Time
	for i := 0; i < len(as.Nodes); i++ {
		notFound = true
		for j := 0; j < len(bs.Nodes); j++ {
			if as.Nodes[i].Uuid == bs.Nodes[j].Uuid {
				notFound = false

				aTime, err = ptypes.Timestamp(as.Nodes[i].LastSeen)
				if err != nil {
					return nil, err
				}

				bTime, err = ptypes.Timestamp(bs.Nodes[j].LastSeen)
				if err != nil {
					return nil, err
				}

				if aTime.After(bTime) {
					newState.Nodes = append(newState.Nodes, as.Nodes[i])
				} else {
					newState.Nodes = append(newState.Nodes, bs.Nodes[j])
				}

				break
			}
		}

		if notFound {
			newState.Nodes = append(newState.Nodes, as.Nodes[i])
		}
	}

	for j := 0; j < len(bs.Nodes); j++ {
		notFound = true
		for i := 0; i < len(as.Nodes); i++ {
			if (as.Nodes)[i].Uuid == (bs.Nodes)[j].Uuid {
				notFound = false
				break
			}
		}

		if notFound {
			newState.Nodes = append(newState.Nodes, bs.Nodes[j])
		}
	}

	return
}
