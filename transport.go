package go_sheep

type Transporter interface {
	Ping(s *State, a string) (*State, error)
	IndirectPing(s *State, a []string, t string) ([]*State, error)
}
