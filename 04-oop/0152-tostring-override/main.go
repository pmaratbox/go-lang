package main

import "fmt"

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("Point(%d, %d)", p.x, p.y)
}

func main() {
	p := Point{1, 2}
	fmt.Println(p)
}
