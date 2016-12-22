package go_sheep

type SWIM interface {
	// Ping sends state 's' to an address:port 'a' and returns State of 'a' OR error
	Ping(s *State, a string) (*State, error)

	// RandomizedTarget returns a Node to ping that is not equal to 'currentNodeInfo'
	// or error if no valid address is found.
	GetRandomizedTarget(s *State, currentNodeInfo *Node) (*Node, error)

	// Checkers returns 'n' addresses:ports to check the existence of 't', which is also an address:port.
	GetCheckers(s *State, t, cur *Node, n int) ([]*Node, error)

	// IndirectPing sends State 's' to the addresses:ports 'd' so that it can resend the state 's' to the target
	// address:port 't'. It's assumed that when calling this function 't' an initial Ping attempt has failed. Three
	// scenarios can be returned:
	// - An array of states where every delegated node and the target has returned their state
	// - An array of states where just the delegated nodes has returned their state
	// - An array of states with a single state from one of the delegated nodes.
	// - No answer from any node.
	IndirectPing(s *State, d []string, t string) ([]*State, error)

	//CheckNode must be triggered by a remote node 'source' to check target 't' passing state 'a'.
	CheckNode(s *State, t, source string)
}
