package main

import (
	"fmt"

	"go.uber.org/dig"
)

// Animal is the abstraction we resolve by.
type Animal interface {
	Sound() string
}

// Dog is the concrete implementation bound to Animal.
type Dog struct{}

func (d *Dog) Sound() string { return "woof" }

// NewAnimal binds the Animal interface to a *Dog implementation.
// The return type Animal is what dig records in the graph, so the
// service is resolved by the interface, not by the concrete type.
func NewAnimal() Animal { return &Dog{} }

func main() {
	c := dig.New()

	// Register the Dog implementation under the Animal interface.
	if err := c.Provide(NewAnimal); err != nil {
		panic(err)
	}

	// Resolve by the interface and call its method.
	if err := c.Invoke(func(a Animal) {
		fmt.Println(a.Sound())
	}); err != nil {
		panic(err)
	}
}
