package main

import "fmt"

type Shape interface {
	area() int
}

type Rect struct {
	w, h int
}

func (r Rect) area() int {
	return r.w * r.h
}

type Square struct {
	s int
}

func (sq Square) area() int {
	return sq.s * sq.s
}

func main() {
	shapes := []Shape{Rect{2, 3}, Square{4}}
	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
