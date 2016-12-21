package go_sheep

import (
	"errors"
	"time"
)

type MembershipMessage struct {
	UUID     string
	Address  string
	LastSeen time.Time
}

type State []MembershipMessage

func MergeState(as, bs *State) (newState *State) {
	newState = &State{}
	*newState = make([]MembershipMessage, 0)

	notFound := true
	var aTime, bTime time.Time
	for i := 0; i < len(*as); i++ {
		notFound = true
		for j := 0; j < len(*bs); j++ {
			if (*as)[i].UUID == (*bs)[j].UUID {
				notFound = false
				aTime = (*as)[i].LastSeen
				if (*as)[i].LastSeen.After((*bs)[j].LastSeen) {
					(*newState) = append((*newState), (*as)[i])
				} else {
					(*newState) = append((*newState), (*bs)[j])
				}

				break
			}
		}

		if notFound {
			(*newState) = append((*newState), (*as)[i])
		}
	}

	for j := 0; j < len(*bs); j++ {
		notFound = true
		for i := 0; i < len(*as); i++ {
			if (*as)[i].UUID == (*bs)[j].UUID {
				notFound = false
				break
			}
		}

		if notFound {
			(*newState) = append((*newState), (*bs)[j])
		}
	}

	return
}

func (s *State) Message(index int) (*MembershipMessage, error) {
	if len(*s) <= index {
		return nil, errors.New("Index out of bounds")
	}

	return &(*s)[index], nil
}
