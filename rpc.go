package go_sheep

import "net/rpc"

type RPCTransport struct {
	rpc.Client
}

func (r *RPCTransport) SendMessage(n Node, m Message) error {
	return nil
}
