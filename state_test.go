package go_sheep

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
)

var initial *State = &State{
	{UUID: "1", Address: "1:1234"},
	{UUID: "2", Address: "2:2234"},
	{UUID: "3", Address: "3:9876"},
}

var arrivedMessage *State = &State{
	{UUID: "1", Address: "1:1234"},
	{UUID: "2", Address: "2:22234"},
	{UUID: "3", Address: "6:3334"},
	{UUID: "4", Address: "4:3334"},
}

func TestState_Merge(t *testing.T) {
	(*initial)[0].LastSeen, _ = ptypes.TimestampProto(time.Now())
	(*initial)[1].LastSeen, _ = ptypes.TimestampProto(time.Now())
	(*initial)[2].LastSeen, _ = ptypes.TimestampProto(time.Now().Add(time.Second * 5))

	(*arrivedMessage)[0].LastSeen, _ = ptypes.TimestampProto(time.Now())
	(*arrivedMessage)[1].LastSeen, _ = ptypes.TimestampProto(time.Now().Add(time.Second * 10))
	(*arrivedMessage)[2].LastSeen, _ = ptypes.TimestampProto(time.Now())
	(*arrivedMessage)[3].LastSeen, _ = ptypes.TimestampProto(time.Now().Add(-(time.Second * 10)))

	original, err := MergeState(initial, arrivedMessage)
	if err != nil {
		t.Fatal(err)
	}

	if len(*original) != len(*arrivedMessage) {
		t.Error()
	}

	if (*original)[0].Address != (*initial)[0].Address {
		t.Error()
	}

	if (*original)[1].Address != (*arrivedMessage)[1].Address {
		t.Error()
	}

	if (*original)[2].Address != (*initial)[2].Address {
		t.Error()
	}

	if (*original)[3] != (*arrivedMessage)[3] {
		t.Error()
	}

	original, err = MergeState(arrivedMessage, initial)
	if err != nil {
		t.Fatal(err)
	}

	if len(*original) != len(*arrivedMessage) {
		t.Error()
	}

	if (*original)[0].Address != (*initial)[0].Address {
		t.Error()
	}

	if (*original)[1].Address != (*arrivedMessage)[1].Address {
		t.Error()
	}

	if (*original)[2].Address != (*initial)[2].Address {
		t.Error()
	}

	if (*original)[3] != (*arrivedMessage)[3] {
		t.Error()
	}
}
