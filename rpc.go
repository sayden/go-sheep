package go_sheep

import "net/rpc"

type RPCTransport struct {
	rpc.Client
}
