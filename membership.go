package go_sheep

import (
	"errors"
	"time"

	"math/rand"

	"github.com/golang/protobuf/ptypes"
)

func newMembershipError(err ...error) error {
	return &Error{
		File:    "membership.go",
		Method:  "RandomizedTarget",
		Package: "go_sheep",
		Errs:    err,
	}
}

func GetRandomizedTarget(s *State, currentNodeInfo *Node) (n *Node, err error) {
	if len(s.Nodes) < 2 {
		return nil, newMembershipError(errors.New("Only one node to randomize"))
	}

	rand.Seed(time.Now().UnixNano())

	for {
		n = s.Nodes[rand.Intn(len(s.Nodes))]
		if n != currentNodeInfo {
			return
		}
	}

	err = newMembershipError(errors.New("Could not find valid randomized target"))

	return
}

func MergeState(as, bs *State) (newState *State, err error) {
	newState = &State{
		Nodes: make([]*Node, 0),
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

func GetCheckers(s *State, targetNode *Node, currentNode *Node, n int) (foundNodes []*Node, err error) {
	if len(s.Nodes) < n+1 {
		return nil, newMembershipError(errors.New("Not enough nodes in state"))
	}

	rand.Seed(time.Now().UnixNano())

	foundNodes = make([]*Node, n)
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

	return
}
