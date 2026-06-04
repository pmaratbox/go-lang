package main

import (
	"fmt"
	"strings"
)

func countingSort(a []int) []int {
	max := 0
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	count := make([]int, max+1)
	for _, v := range a {
		count[v]++
	}
	out := make([]int, 0, len(a))
	for v := 0; v <= max; v++ {
		for c := 0; c < count[v]; c++ {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	a := []int{3, 1, 2, 3, 1}
	sorted := countingSort(a)
	parts := make([]string, len(sorted))
	for i, v := range sorted {
		parts[i] = fmt.Sprint(v)
	}
	fmt.Println(strings.Join(parts, " "))
}
