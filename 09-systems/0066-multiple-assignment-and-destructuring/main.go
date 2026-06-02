package main

import "fmt"

func main() {
	a, b := 1, 2
	a, b = b, a
	fmt.Printf("%d %d\n", a, b)

	pair := [2]int{3, 4}
	x, y := pair[0], pair[1]
	fmt.Printf("%d %d\n", x, y)
}
