package main

import "fmt"

type Point struct{ x, y int }

func main() {
	def := Point{}
	over := Point{x: 5}
	fmt.Println(def.x, def.y)
	fmt.Println(over.x, over.y)
}
