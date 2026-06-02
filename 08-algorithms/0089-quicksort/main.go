package main

import (
	"fmt"
	"strings"
)

func quicksort(items []int) []int {
	if len(items) <= 1 {
		return items
	}
	pivot := items[0]
	var less, greater []int
	for _, x := range items[1:] {
		if x < pivot {
			less = append(less, x)
		} else {
			greater = append(greater, x)
		}
	}
	result := quicksort(less)
	result = append(result, pivot)
	result = append(result, quicksort(greater)...)
	return result
}

func main() {
	data := []int{3, 1, 4, 1, 5, 2}
	sorted := quicksort(data)
	parts := []string{}
	for _, x := range sorted {
		parts = append(parts, fmt.Sprint(x))
	}
	fmt.Println(strings.Join(parts, " "))
}
