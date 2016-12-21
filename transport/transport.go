package transport

import "github.com/sayden/go-sheep"

type Transporter interface {
	Ping(s *go_sheep.State, a string) (*go_sheep.State, error)
}
