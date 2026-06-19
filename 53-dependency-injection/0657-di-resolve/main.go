package main

import (
	"fmt"

	"go.uber.org/dig"
)

// Greeter is the service we register and resolve.
type Greeter struct{}

func (g *Greeter) Greet() string { return "hello" }

// NewGreeter is the constructor dig uses to build the service.
func NewGreeter() *Greeter { return &Greeter{} }

func main() {
	c := dig.New()

	// Register the Greeter service with the container.
	if err := c.Provide(NewGreeter); err != nil {
		panic(err)
	}

	// Resolve the service from the container and call its method.
	if err := c.Invoke(func(g *Greeter) {
		fmt.Println(g.Greet())
	}); err != nil {
		panic(err)
	}
}
