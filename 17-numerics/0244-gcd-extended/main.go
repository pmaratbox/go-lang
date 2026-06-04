package main

import "fmt"

func extGCD(a, b int) (g, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := extGCD(b, a%b)
	return g, y1, x1 - (a/b)*y1
}

func main() {
	g, x, y := extGCD(30, 12)
	fmt.Printf("%d %d %d\n", g, x, y)
}
