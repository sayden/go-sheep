package go_sheep

type SWIM interface {
	// Ping sends state 's' to an address:port 'a' and returns State of 'a' or error
	SWIMPing(s *State, a string) (*State, error)

	// RandomizedTarget returns an address:port to ping or error if no valid address is found.
	RandomizedTarget(p *State) (*string, error)

	// Checkers returns 'n' addresses:ports to check the existence of 't', which is also an address:port.
	Checkers(t string, n int) ([]*string, error)

	// Check sends State 's' to the address:port 'a' so that it can resend the state 's' to the target
	// address:port 't'. It's assumed that when calling CheckNode 't' an initial Ping attempt has failed. It must
	// return an array of states, on the best case of 't' and 'a', 't' only or in the worst case error
	CheckNode(s *State, a []string, t string) ([]*State, error)
}


