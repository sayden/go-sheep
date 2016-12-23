package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sayden/go-sheep"
)

var state go_sheep.State

var loopTime int

var currentNode go_sheep.Node = go_sheep.Node{
	Address:  fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]),
	Uuid:     getNewUUID(),
	LastSeen: time.Now(),
}

func main() {
	state.Nodes = []go_sheep.Node{currentNode}

	//init listener

	//launch loop
	loop()
}

func loop() {
	for {
		//Take a new peer to check

		//Wait n seconds
		time.Sleep(time.Second * loopTime)
	}
}

func getNewUUID() string {
	return ""
}
