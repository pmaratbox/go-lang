package main

import "fmt"

type mediator struct {
	a *colleague
	b *colleague
}

func (m *mediator) send(from *colleague, msg string) {
	if from == m.a {
		m.b.receive(msg)
	} else {
		m.a.receive(msg)
	}
}

type colleague struct {
	name string
	med  *mediator
}

func (c *colleague) send(msg string) { c.med.send(c, msg) }
func (c *colleague) receive(msg string) {
	fmt.Printf("%s got: %s\n", c.name, msg)
}

func main() {
	m := &mediator{}
	a := &colleague{name: "A", med: m}
	b := &colleague{name: "B", med: m}
	m.a, m.b = a, b
	a.send("hi")
}
