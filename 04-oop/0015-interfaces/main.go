package main

import "fmt"

type Shape interface {
	name() string
	area() int
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) name() string { return "rectangle" }
func (r Rectangle) area() int    { return r.width * r.height }

type Square struct {
	side int
}

func (s Square) name() string { return "square" }
func (s Square) area() int    { return s.side * s.side }

func main() {
	shapes := []Shape{
		Rectangle{width: 3, height: 4},
		Square{side: 5},
	}
	for _, s := range shapes {
		fmt.Printf("%s area: %d\n", s.name(), s.area())
	}
}
