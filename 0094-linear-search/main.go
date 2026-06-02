package main

import "fmt"

func main() {
	nums := []int{4, 2, 7, 1, 9}
	target := 7
	index := -1
	for i, n := range nums {
		if n == target {
			index = i
			break
		}
	}
	fmt.Printf("found %d at index %d\n", target, index)
}
