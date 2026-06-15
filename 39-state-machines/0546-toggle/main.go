package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/looplab/fsm"
)

func main() {
	f := fsm.NewFSM(
		"off",
		fsm.Events{
			{Name: "toggle", Src: []string{"off"}, Dst: "on"},
			{Name: "toggle", Src: []string{"on"}, Dst: "off"},
		},
		nil,
	)

	// Fire the fixed event sequence: off -> on -> off -> on.
	for i := 0; i < 3; i++ {
		f.Event(context.Background(), "toggle")
	}

	fmt.Println(strings.ToLower(f.Current()))
}
