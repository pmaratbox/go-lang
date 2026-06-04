package main

import "fmt"

type coffee interface {
	cost() int
}

type base struct{}

func (base) cost() int { return 2 }

type milk struct{ inner coffee }

func (m milk) cost() int { return m.inner.cost() + 1 }

type sugar struct{ inner coffee }

func (s sugar) cost() int { return s.inner.cost() + 1 }

func main() {
	var c coffee = base{}
	c = milk{inner: c}
	c = sugar{inner: c}
	fmt.Println(c.cost())
}
