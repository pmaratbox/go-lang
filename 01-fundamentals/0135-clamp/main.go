package main

import "fmt"

func clamp(x, lo, hi int) int {
	return max(lo, min(x, hi))
}

func main() {
	fmt.Printf("%d %d\n", clamp(15, 0, 10), clamp(-3, 0, 10))
}
