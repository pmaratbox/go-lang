package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/looplab/fsm"
)

func main() {
	f := fsm.NewFSM(
		"idle",
		fsm.Events{
			{Name: "submit", Src: []string{"idle"}, Dst: "pending"},
			{Name: "approve", Src: []string{"pending"}, Dst: "approved"},
		},
		nil,
	)

	f.Event(context.Background(), "submit")
	f.Event(context.Background(), "approve")

	fmt.Println(strings.ToLower(f.Current()))
}
