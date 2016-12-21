package main

import (
	"time"
	"github.com/sayden/go-sheep"
)

type Node struct {
	IP   string
	Port int
}

type Message interface {
	GetPayload() []byte
}

type Transport interface {
	SendMessage(Node, Message) error
}

var initials *go_sheep.State = &go_sheep.State{
	{
		UUID:     "1",
		Address:       "1",
		Port:     1234,
		LastSeen: time.Now(),
	},
	{
		UUID:     "2",
		Address:       "2",
		Port:     2234,
		LastSeen: time.Now(),
	},
	{
		UUID:     "3",
		Address:       "3",
		Port:     9876,
		LastSeen: time.Now().Add(time.Second * 5),
	},
}

var arrivedMessages *go_sheep.State = &go_sheep.State{
	{
		UUID:     "1",
		Address:       "1",
		Port:     1234,
		LastSeen: time.Now(),
	},
	{
		UUID:     "2",
		Address:       "2",
		Port:     22234,
		LastSeen: time.Now().Add(time.Second * 10),
	},
	{
		UUID:     "3",
		Address:       "6",
		Port:     3334,
		LastSeen: time.Now(),
	},
	{
		UUID:     "4",
		Address:       "4",
		Port:     3334,
		LastSeen: time.Now().Add(-(time.Second * 10)),
	},
}

func main() {
	go_sheep.MergeState(initials, arrivedMessages)
}
