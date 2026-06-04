package main

import "fmt"

func sign(x int) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}

func main() {
	fmt.Printf("%d %d %d\n", sign(-5), sign(0), sign(5))
}
