package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	fmt.Printf("point: (%d, %d)\n", p1.X, p1.Y)
	equal := "no"
	if p1 == p2 {
		equal = "yes"
	}
	fmt.Println("equal:", equal)
}
