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
			{Name: "start", Src: []string{"idle"}, Dst: "running"},
			{Name: "reset", Src: []string{"running"}, Dst: "idle"},
		},
		nil,
	)

	f.Event(context.Background(), "start")
	f.Event(context.Background(), "reset")

	fmt.Println(strings.ToLower(f.Current()))
}
