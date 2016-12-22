package go_sheep

import "github.com/sayden/go-sheep"

type Transporter interface {
	Ping(s *go_sheep.State, a string) (*go_sheep.State, error)
	IndirectPing(s *go_sheep.State, a string, t string) ([]*go_sheep.State, error)
}
