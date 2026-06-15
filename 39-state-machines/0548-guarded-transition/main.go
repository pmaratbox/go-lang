package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/looplab/fsm"
)

func main() {
	// Door FSM with a guarded transition: the "open" event is ONLY valid from
	// the "unlocked" state. Each event's Src list defines the guard — firing an
	// event from a state not listed in Src is rejected and leaves the state
	// unchanged.
	f := fsm.NewFSM(
		"locked",
		fsm.Events{
			{Name: "unlock", Src: []string{"locked"}, Dst: "unlocked"},
			{Name: "open", Src: []string{"unlocked"}, Dst: "open"},
		},
		fsm.Callbacks{},
	)

	// Fire the fixed event sequence. "open" is guarded to unlocked, so it only
	// succeeds because we unlock first.
	f.Event(context.Background(), "unlock")
	f.Event(context.Background(), "open")

	// Print the resulting state, lowercased.
	fmt.Println(strings.ToLower(f.Current()))
}
