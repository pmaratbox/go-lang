package main

import "fmt"

func first[T any](items []T) T {
	return items[0]
}

func main() {
	ints := []int{1, 2, 3}
	strs := []string{"a", "b", "c"}

	fmt.Printf("first int: %d\n", first(ints))
	fmt.Printf("first string: %s\n", first(strs))
}
