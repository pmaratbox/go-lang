package main

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

func main() {
	count := 0

	// A simple cyclic light: green -> yellow -> red -> green.
	f := fsm.NewFSM(
		"green",
		fsm.Events{
			{Name: "next", Src: []string{"green"}, Dst: "yellow"},
			{Name: "next", Src: []string{"yellow"}, Dst: "red"},
			{Name: "next", Src: []string{"red"}, Dst: "green"},
		},
		fsm.Callbacks{
			// Runs after every successful event, i.e. on each transition.
			"after_event": func(_ context.Context, _ *fsm.Event) {
				count++
			},
		},
	)

	// Fire a fixed sequence of 3 valid events.
	f.Event(context.Background(), "next")
	f.Event(context.Background(), "next")
	f.Event(context.Background(), "next")

	// The counter comes from the per-transition action, not a literal.
	fmt.Println(count)
}
