package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p1 := Point{1, 2}
	p2 := p1
	p2.X = 9
	fmt.Printf("original: (%d, %d)\n", p1.X, p1.Y)
	fmt.Printf("updated: (%d, %d)\n", p2.X, p2.Y)
}
