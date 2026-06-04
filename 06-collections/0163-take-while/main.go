package main

import (
	"fmt"
	"strings"
)

func takeWhile(xs []int, pred func(int) bool) []int {
	var out []int
	for _, x := range xs {
		if !pred(x) {
			break
		}
		out = append(out, x)
	}
	return out
}

func main() {
	xs := []int{1, 2, 3, 4, 1}
	taken := takeWhile(xs, func(x int) bool { return x < 3 })
	parts := make([]string, len(taken))
	for i, v := range taken {
		parts[i] = fmt.Sprintf("%d", v)
	}
	fmt.Println(strings.Join(parts, " "))
}
