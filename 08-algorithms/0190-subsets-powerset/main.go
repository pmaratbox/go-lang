package main

import (
	"fmt"
	"strings"
)

func main() {
	elems := []int{1, 2, 3}
	n := len(elems)
	var lines []string
	for mask := 0; mask < 1<<n; mask++ {
		var parts []string
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				parts = append(parts, fmt.Sprint(elems[i]))
			}
		}
		if len(parts) == 0 {
			lines = append(lines, "{}")
		} else {
			lines = append(lines, strings.Join(parts, " "))
		}
	}
	fmt.Println(strings.Join(lines, "\n"))
}
