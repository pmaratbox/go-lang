package main

import (
	"fmt"

	"go.uber.org/dig"
)

// A is the first dependency. x() returns "a".
type A struct{}

func (a *A) X() string { return "a" }

// B is the second dependency. y() returns "b".
type B struct{}

func (b *B) Y() string { return "b" }

// Service depends on BOTH A and B. run() returns x()+y().
type Service struct {
	a *A
	b *B
}

func (s *Service) Run() string { return s.a.X() + s.b.Y() }

// Constructors dig uses to build each piece of the graph.
func NewA() *A { return &A{} }
func NewB() *B { return &B{} }

// NewService takes both dependencies; dig wires them automatically.
func NewService(a *A, b *B) *Service { return &Service{a: a, b: b} }

func main() {
	c := dig.New()

	// Register all three constructors with the container.
	for _, provide := range []interface{}{NewA, NewB, NewService} {
		if err := c.Provide(provide); err != nil {
			panic(err)
		}
	}

	// Resolve Service; dig builds A and B and injects both, then run().
	if err := c.Invoke(func(s *Service) {
		fmt.Println(s.Run())
	}); err != nil {
		panic(err)
	}
}
