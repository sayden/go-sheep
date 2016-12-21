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
