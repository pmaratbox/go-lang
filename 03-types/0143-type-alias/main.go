package main

import "fmt"

// Meters is a defined type backed by int, giving the value semantic meaning.
type Meters int

func main() {
	var distance Meters = 5
	fmt.Printf("distance: %d\n", distance)
}
