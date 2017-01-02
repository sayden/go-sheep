package transport

import (
	"github.com/sayden/go-sheep"
	"github.com/uber-go/zap"
)

type Server interface {
	ServeJoin(*go_sheep.Node) (*go_sheep.State, error)
	ServePing(*go_sheep.State) (*go_sheep.State, error)
	ServeIndirectCheck(*go_sheep.DelegateCheckRequest) (*go_sheep.State, error)
	Serve(string, chan struct{}, zap.Logger) error
}
