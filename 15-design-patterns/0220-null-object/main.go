package main

import "fmt"

type logger interface {
	log()
}

type nullLogger struct{}

func (nullLogger) log() {}

type realLogger struct{ count int }

func (r *realLogger) log() { r.count++ }

func main() {
	var nl logger = nullLogger{}
	rl := &realLogger{}
	nl.log()
	rl.log()
	fmt.Println(rl.count)
}
