package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	sum := 0
	for _, n := range nums {
		if n%2 == 0 {
			sum += n * 2
		}
	}

	fmt.Printf("sum: %d\n", sum)
}
