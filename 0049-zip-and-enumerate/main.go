package main

import (
	"fmt"
	"strings"
)

func main() {
	letters := []string{"a", "b", "c"}
	nums := []int{1, 2, 3}
	parts := make([]string, len(letters))
	for i, k := range letters {
		parts[i] = fmt.Sprintf("%s=%d", k, nums[i])
	}
	fmt.Println(strings.Join(parts, " "))
}
