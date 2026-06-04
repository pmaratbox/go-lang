package main

import "fmt"

type color interface {
	name() string
}

type red struct{}

func (red) name() string { return "red" }

type circle struct{ c color }

func (s circle) describe() string { return s.c.name() + " circle" }

func main() {
	s := circle{c: red{}}
	fmt.Println(s.describe())
}
