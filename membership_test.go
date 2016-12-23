package go_sheep

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
)

var initial *State = &State{
	Nodes: []*Node{
		{Uuid: "1", Address: "1:1234"},
		{Uuid: "2", Address: "2:2234"},
		{Uuid: "3", Address: "3:9876"},
	},
}

var arrivedMessage *State = &State{
	Nodes: []*Node{
		{Uuid: "1", Address: "1:1234"},
		{Uuid: "2", Address: "2:22234"},
		{Uuid: "3", Address: "6:3334"},
		{Uuid: "4", Address: "4:3334"},
	},
}

func TestState_Merge(t *testing.T) {
	initial.Nodes[0].LastSeen, _ = ptypes.TimestampProto(time.Now())
	initial.Nodes[1].LastSeen, _ = ptypes.TimestampProto(time.Now())
	initial.Nodes[2].LastSeen, _ = ptypes.TimestampProto(time.Now().Add(time.Second * 5))

	arrivedMessage.Nodes[0].LastSeen, _ = ptypes.TimestampProto(time.Now())
	arrivedMessage.Nodes[1].LastSeen, _ = ptypes.TimestampProto(time.Now().Add(time.Second * 10))
	arrivedMessage.Nodes[2].LastSeen, _ = ptypes.TimestampProto(time.Now())
	arrivedMessage.Nodes[3].LastSeen, _ = ptypes.TimestampProto(time.Now().Add(-(time.Second * 10)))

	original, err := MergeState(initial, arrivedMessage)
	if err != nil {
		t.Fatal(err)
	}

	if len(original.Nodes) != len(arrivedMessage.Nodes) {
		t.Error()
	}

	if original.Nodes[0].Address != initial.Nodes[0].Address {
		t.Error()
	}

	if original.Nodes[1].Address != arrivedMessage.Nodes[1].Address {
		t.Error()
	}

	if original.Nodes[2].Address != initial.Nodes[2].Address {
		t.Error()
	}

	if original.Nodes[3] != arrivedMessage.Nodes[3] {
		t.Error()
	}

	original, err = MergeState(arrivedMessage, initial)
	if err != nil {
		t.Fatal(err)
	}

	if len(original.Nodes) != len(arrivedMessage.Nodes) {
		t.Error()
	}

	if original.Nodes[0].Address != initial.Nodes[0].Address {
		t.Error()
	}

	if original.Nodes[1].Address != arrivedMessage.Nodes[1].Address {
		t.Error()
	}

	if original.Nodes[2].Address != initial.Nodes[2].Address {
		t.Error()
	}

	if original.Nodes[3] != arrivedMessage.Nodes[3] {
		t.Error()
	}
}

func TestGetRandomizedTarget(t *testing.T) {
	currentNode := &Node{
		Address: "12345:8080",
		Uuid:    "1",
	}

	state := &State{
		Nodes: []*Node{
			{
				Address: "234234:12312",
				Uuid:    "2",
			},
			currentNode,
			{
				Address: "56346:213123",
				Uuid:    "3",
			},
		},
	}

	for i := 0; i < 100; i++ {
		node, err := GetRandomizedTarget(state, currentNode)
		if err != nil {
			t.Fatal(err)
		}

		if node == currentNode {
			t.Fail()
		}
	}

	state.Nodes = state.Nodes[:1]
	node, err := GetRandomizedTarget(state, currentNode)
	if err == nil {
		t.Fail()
	}

	if node != nil {
		t.Fail()
	}

	state.Nodes = state.Nodes[:1]
	node, err = GetRandomizedTarget(state, currentNode)
	if err == nil {
		t.Fail()
	}

	if node != nil {
		t.Fail()
	}
}

func TestGetCheckers(t *testing.T) {
	currentNode := &Node{
		Address: "23234:1233",
		Uuid:    "1",
	}

	state := &State{
		Nodes: []*Node{
			{
				Address: "23213:459345",
				Uuid:    "2",
			},
			currentNode,
		},
	}

	target, err := GetRandomizedTarget(state, currentNode)
	if err != nil {
		t.Fatal(err)
	}

	_, err = GetCheckers(state, target, currentNode, 2)
	if err == nil {
		t.Fail()
	}

	state.Nodes = append(state.Nodes, &Node{
		Address: "1234:0239",
		Uuid:    "3",
	},&Node{
		Address:"45243:123",
		Uuid:"4",
	})

	nodes, err := GetCheckers(state, target, currentNode, 2)
	if err != nil {
		t.Fatal(err)
	}

	if len(nodes) != 2 {
		t.Error("Not enough nodes returned")
	}

	for _, node := range nodes {
		if node == currentNode || node == target{
			t.Fail()
		}
	}
}
