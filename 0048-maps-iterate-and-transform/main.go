package main

import "fmt"

func main() {
	scores := map[string]int{"a": 1, "b": 2, "c": 3}
	total := 0
	for _, v := range scores {
		total += v
	}
	fmt.Printf("sum: %d\n", total)
}
