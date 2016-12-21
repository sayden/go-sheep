package go_sheep

import (
	"testing"
	"time"
	"github.com/kr/pretty"
)

var initial *State = &State{
	{
		UUID:     "1",
		Address:  "1:1234",
		LastSeen: time.Now(),
	},
	{
		UUID:     "2",
		Address:  "2:2234",
		LastSeen: time.Now(),
	},
	{
		UUID:     "3",
		Address:  "3:9876",
		LastSeen: time.Now().Add(time.Second * 5),
	},
}

var arrivedMessage *State = &State{
	{
		UUID:     "1",
		Address:  "1:1234",
		LastSeen: time.Now(),
	},
	{
		UUID:     "2",
		Address:  "2:22234",
		LastSeen: time.Now().Add(time.Second * 10),
	},
	{
		UUID:     "3",
		Address:  "6:3334",
		LastSeen: time.Now(),
	},
	{
		UUID:     "4",
		Address:  "4:3334",
		LastSeen: time.Now().Add(-(time.Second * 10)),
	},
}

func TestState_Merge(t *testing.T) {
	original := MergeState(initial, arrivedMessage)
	//t.Logf("%# v", pretty.Formatter(original))

		defer func(){
			if err := recover(); err != nil {
				t.Logf("%# v\nERROR: %# v", pretty.Formatter(original), pretty.Formatter(err))
				t.Fail()
			}
		}()

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

	original = MergeState(arrivedMessage, initial)

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
