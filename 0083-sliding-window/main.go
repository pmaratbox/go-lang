package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []int{1, 2, 3, 4}
	size := 2
	for i := 0; i+size <= len(nums); i++ {
		parts := []string{}
		for _, x := range nums[i : i+size] {
			parts = append(parts, fmt.Sprint(x))
		}
		fmt.Println(strings.Join(parts, " "))
	}
}
