package main

import "fmt"

type Shape interface {
	area() int
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) area() int {
	return r.width * r.height
}

type Triangle struct {
	base, height int
}

func (t Triangle) area() int {
	return t.base * t.height / 2
}

func main() {
	shapes := []Shape{
		Rectangle{width: 2, height: 3},
		Triangle{base: 4, height: 4},
	}

	total := 0
	for _, s := range shapes {
		total += s.area()
	}
	fmt.Printf("total area: %d\n", total)
}
