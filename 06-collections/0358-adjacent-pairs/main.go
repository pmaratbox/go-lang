package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []int{1, 2, 3, 4}
	parts := make([]string, 0, len(nums)-1)
	for i := 0; i+1 < len(nums); i++ {
		parts = append(parts, fmt.Sprintf("%d,%d", nums[i], nums[i+1]))
	}
	fmt.Println(strings.Join(parts, " "))
}
