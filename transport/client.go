package transport

import "github.com/sayden/go-sheep"

type Client interface {
	DoJoin(in, targetServer *go_sheep.Node) (state *go_sheep.State, err error)
	DoPing(s *go_sheep.State, a string) (state *go_sheep.State, err error)
	DoIndirectPing(s *go_sheep.State, delegatedNodes []*go_sheep.Node, t *go_sheep.Node) (states []*go_sheep.State, err error)
}