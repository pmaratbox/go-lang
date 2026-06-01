package main

import "fmt"

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func main() {
	lo, hi := minMax(3, 7)
	fmt.Printf("min: %d\n", lo)
	fmt.Printf("max: %d\n", hi)
}
