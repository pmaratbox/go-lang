package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	rotated := append([]int{}, nums[k:]...)
	rotated = append(rotated, nums[:k]...)
	parts := []string{}
	for _, x := range rotated {
		parts = append(parts, fmt.Sprint(x))
	}
	fmt.Println(strings.Join(parts, " "))
}
