package main

import "fmt"

type shape interface {
	kind() string
}

type circle struct{}

func (circle) kind() string { return "circle" }

type square struct{}

func (square) kind() string { return "square" }

func newShape(name string) shape {
	switch name {
	case "circle":
		return circle{}
	case "square":
		return square{}
	}
	return nil
}

func main() {
	a := newShape("circle")
	b := newShape("square")
	fmt.Println(a.kind(), b.kind())
}
