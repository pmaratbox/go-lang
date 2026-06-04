package main

import (
	"fmt"
	"strings"
)

func reverse(xs []int) []int {
	if len(xs) == 0 {
		return nil
	}
	return append(reverse(xs[1:]), xs[0])
}

func main() {
	r := reverse([]int{1, 2, 3})
	parts := make([]string, len(r))
	for i, x := range r {
		parts[i] = fmt.Sprint(x)
	}
	fmt.Println(strings.Join(parts, " "))
}
