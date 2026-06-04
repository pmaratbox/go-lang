package main

import (
	"fmt"
	"strings"
)

type stepper interface {
	step() string
}

func run(s stepper) string {
	return strings.Join([]string{"start", s.step(), "end"}, " ")
}

type worker struct{}

func (worker) step() string { return "work" }

func main() {
	fmt.Println(run(worker{}))
}
