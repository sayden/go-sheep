package failure_detector

import "github.com/sayden/go-sheep"

type swim struct {
	go_sheep.SWIM
	Transport go_sheep.Transporter
}

func (swim *swim) SWIMPing(s *go_sheep.State, a string) (*go_sheep.State, error) {
	return swim.Transport.Ping(s, a)
}

func (swim *swim) RandomizedTarget(p *go_sheep.State) (*string, error) {
	panic("not implemented")
}

func (swim *swim) Checkers(t string, n int) ([]*string, error) {
	panic("not implemented")
}

func (swim *swim) IndirectPing(s *go_sheep.State, a []string, t string) ([]*go_sheep.State, error) {
	panic("not implemented")
}
