package main

import (
	"fmt"
	"strings"
)

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func mergeSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}
	mid := len(items) / 2
	return merge(mergeSort(items[:mid]), mergeSort(items[mid:]))
}

func main() {
	data := []int{3, 1, 4, 1, 5, 2}
	sorted := mergeSort(data)
	parts := []string{}
	for _, x := range sorted {
		parts = append(parts, fmt.Sprint(x))
	}
	fmt.Println(strings.Join(parts, " "))
}
