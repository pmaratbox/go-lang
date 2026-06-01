package main

import "fmt"

func word(n int) string {
	switch n {
	case 1:
		return "one"
	case 2:
		return "two"
	default:
		return "many"
	}
}

func main() {
	for _, n := range []int{1, 2, 5} {
		fmt.Printf("%d: %s\n", n, word(n))
	}
}
