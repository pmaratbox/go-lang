package main

import "fmt"

type Calc struct {
	value int
}

func (c *Calc) Add(n int) *Calc {
	c.value += n
	return c
}

func (c *Calc) Multiply(n int) *Calc {
	c.value *= n
	return c
}

func main() {
	c := &Calc{value: 5}
	fmt.Println(c.Add(3).Multiply(2).value)
}
