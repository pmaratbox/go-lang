package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("Point(x=%d, y=%d)", p.X, p.Y)
}

func main() {
	p := Point{X: 1, Y: 2}
	fmt.Println(p)
}
