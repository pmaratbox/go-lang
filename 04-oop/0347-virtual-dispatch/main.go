package main

import (
	"fmt"
	"strings"
)

type Shape interface {
	describe() string
}

type Circle struct{}

func (Circle) describe() string { return "circle" }

type Square struct{}

func (Square) describe() string { return "square" }

type Triangle struct{}

func (Triangle) describe() string { return "triangle" }

func main() {
	shapes := []Shape{Circle{}, Square{}, Triangle{}}
	parts := make([]string, len(shapes))
	for i, s := range shapes {
		parts[i] = s.describe()
	}
	fmt.Println(strings.Join(parts, " "))
}
