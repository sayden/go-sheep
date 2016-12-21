package go_sheep

import (
	"time"

	"github.com/golang/protobuf/ptypes"
)

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