package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []int{1, 2, 3, 4}
	sums := make([]int, len(nums))
	total := 0
	for i, n := range nums {
		total += n
		sums[i] = total
	}

	parts := []string{}
	for _, x := range sums {
		parts = append(parts, fmt.Sprint(x))
	}
	fmt.Println(strings.Join(parts, " "))
}
