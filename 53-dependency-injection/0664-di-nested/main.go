package main

import (
	"fmt"

	"go.uber.org/dig"
)

// A is the bottom of the dependency chain.
type A struct{}

func (a *A) V() string { return "a" }

// B depends on A.
type B struct{ a *A }

func (b *B) V() string { return b.a.V() + "b" }

// C depends on B (which depends on A) — a 3-level chain.
type C struct{ b *B }

func (c *C) V() string { return c.b.V() + "c" }

// Constructors dig wires together to build the graph.
func NewA() *A         { return &A{} }
func NewB(a *A) *B     { return &B{a} }
func NewC(b *B) *C     { return &C{b} }

func main() {
	c := dig.New()

	// Register each level of the chain.
	for _, provide := range []interface{}{NewA, NewB, NewC} {
		if err := c.Provide(provide); err != nil {
			panic(err)
		}
	}

	// Resolve C: dig builds A, then B(A), then C(B) automatically.
	if err := c.Invoke(func(svc *C) {
		fmt.Println(svc.V())
	}); err != nil {
		panic(err)
	}
}
