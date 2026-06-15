package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/looplab/fsm"
)

func main() {
	// Turnstile FSM: locked --coin--> unlocked, unlocked --push--> locked.
	f := fsm.NewFSM(
		"locked",
		fsm.Events{
			{Name: "coin", Src: []string{"locked"}, Dst: "unlocked"},
			{Name: "push", Src: []string{"unlocked"}, Dst: "locked"},
		},
		fsm.Callbacks{},
	)

	// Fire the fixed event sequence.
	f.Event(context.Background(), "coin")

	// Print the resulting state, lowercased.
	fmt.Println(strings.ToLower(f.Current()))
}
