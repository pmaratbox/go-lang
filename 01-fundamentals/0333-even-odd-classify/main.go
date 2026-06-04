package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []int{1, 2, 3, 4}
	labels := make([]string, len(nums))
	for i, n := range nums {
		if n%2 == 0 {
			labels[i] = "even"
		} else {
			labels[i] = "odd"
		}
	}
	fmt.Println(strings.Join(labels, " "))
}
