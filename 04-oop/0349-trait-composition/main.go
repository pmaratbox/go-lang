package main

import "fmt"

// A is one capability printing "a".
type A struct{}

func (A) a() { fmt.Print("a") }

// B is another capability printing "b".
type B struct{}

func (B) b() { fmt.Print("b") }

// Combined mixes in both A and B via embedding.
type Combined struct {
	A
	B
}

func main() {
	c := Combined{}
	c.a()
	fmt.Print(" ")
	c.b()
	fmt.Println()
}
