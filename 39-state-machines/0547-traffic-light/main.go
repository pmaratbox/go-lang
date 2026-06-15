package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/looplab/fsm"
)

func main() {
	// Traffic FSM: red --next--> green --next--> yellow --next--> red.
	f := fsm.NewFSM(
		"red",
		fsm.Events{
			{Name: "next", Src: []string{"red"}, Dst: "green"},
			{Name: "next", Src: []string{"green"}, Dst: "yellow"},
			{Name: "next", Src: []string{"yellow"}, Dst: "red"},
		},
		fsm.Callbacks{},
	)

	// Fire the fixed event sequence: red -> green -> yellow.
	f.Event(context.Background(), "next")
	f.Event(context.Background(), "next")

	// Print the resulting state, lowercased.
	fmt.Println(strings.ToLower(f.Current()))
}
