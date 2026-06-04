package main

import "fmt"

type counter struct{ value int }

type command interface {
	execute()
	undo()
}

type addCommand struct {
	c      *counter
	amount int
}

func (a addCommand) execute() { a.c.value += a.amount }
func (a addCommand) undo()    { a.c.value -= a.amount }

func main() {
	c := &counter{value: 0}
	cmd := addCommand{c: c, amount: 5}
	cmd.execute()
	after := c.value
	cmd.undo()
	fmt.Println(after, c.value)
}
