package failure_detector

import "github.com/sayden/go-sheep"

type swim struct {
	transport go_sheep.Transporter
}

func (swim *swim) Ping(s *go_sheep.State, a string) (*go_sheep.State, error) {
	return swim.transport.Ping(s, a)
}

func (swim *swim) GetRandomizedTarget(s *go_sheep.State, currentNodeInfo *go_sheep.Node) (*go_sheep.Node, error) {
	target, err := go_sheep.GetRandomizedTarget(s, currentNodeInfo)
	if err != nil {
		return nil, err
	}

	newState, err := swim.Ping(s, target.Address)
	if err != nil {
		checkers, err := swim.GetCheckers(s, target, currentNodeInfo, 2)
		if err != nil {
			return nil, err
		}

		state, err := swim.IndirectPing(s, checkers, target)
		if err != nil {
			return nil, err
		}
	}
}

func (s *swim) GetCheckers(s *go_sheep.State, t *go_sheep.Node, cur *go_sheep.Node, n int) ([]*go_sheep.Node, error) {
	panic("not implemented")
}

func (s *swim) IndirectPing(s *go_sheep.State, d []string, t string) ([]*go_sheep.State, error) {
	panic("not implemented")
}

func (s *swim) CheckNode(s *go_sheep.State, t string, source string) {
	panic("not implemented")
}
