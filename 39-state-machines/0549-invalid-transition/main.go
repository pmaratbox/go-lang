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

	// Fire "push" from locked: there is NO transition for it, so the FSM
	// rejects the event with an error. We ignore the error; the state is
	// left unchanged by the machine itself.
	_ = f.Event(context.Background(), "push")

	// Print the resulting state, lowercased — still "locked".
	fmt.Println(strings.ToLower(f.Current()))
}
