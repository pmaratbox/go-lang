package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{10, 20, 30, 40, 50}
	slice := nums[1:4]
	parts := make([]string, len(slice))
	for i, n := range slice {
		parts[i] = strconv.Itoa(n)
	}
	fmt.Println("slice:", strings.Join(parts, " "))
}
