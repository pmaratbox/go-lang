package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	size := 3
	for i := 0; i < len(nums); i += size {
		end := i + size
		if end > len(nums) {
			end = len(nums)
		}
		parts := []string{}
		for _, x := range nums[i:end] {
			parts = append(parts, fmt.Sprint(x))
		}
		fmt.Println(strings.Join(parts, " "))
	}
}
