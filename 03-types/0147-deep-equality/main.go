package main

import "fmt"

type Pair struct {
	A, B int
}

type Nested struct {
	Left  Pair
	Right Pair
}

func main() {
	x := Nested{Left: Pair{1, 2}, Right: Pair{3, 4}}
	y := Nested{Left: Pair{1, 2}, Right: Pair{3, 4}}

	// For comparable structs, == compares every nested field structurally.
	eq := "no"
	if x == y {
		eq = "yes"
	}

	fmt.Printf("equal: %s\n", eq)
}
