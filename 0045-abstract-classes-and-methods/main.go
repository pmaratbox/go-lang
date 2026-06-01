package main

import "fmt"

type Shape interface {
	Area() int
}

func describe(s Shape) string {
	return fmt.Sprintf("area: %d", s.Area())
}

type Square struct {
	side int
}

func (s Square) Area() int {
	return s.side * s.side
}

func main() {
	fmt.Println(describe(Square{3}))
}
