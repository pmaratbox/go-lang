package main

import "fmt"

func main() {
	nums := map[string]int{"one": 1, "two": 2, "three": 3}

	fmt.Printf("two: %d\n", nums["two"])
	fmt.Printf("size: %d\n", len(nums))
}
