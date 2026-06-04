package main

import (
	"fmt"
	"strings"
)

func permute(prefix, rest []int, out *[]string) {
	if len(rest) == 0 {
		parts := make([]string, len(prefix))
		for i, v := range prefix {
			parts[i] = fmt.Sprint(v)
		}
		*out = append(*out, strings.Join(parts, " "))
		return
	}
	for i := range rest {
		next := make([]int, 0, len(rest)-1)
		next = append(next, rest[:i]...)
		next = append(next, rest[i+1:]...)
		np := make([]int, 0, len(prefix)+1)
		np = append(np, prefix...)
		np = append(np, rest[i])
		permute(np, next, out)
	}
}

func main() {
	var out []string
	permute(nil, []int{1, 2, 3}, &out)
	fmt.Println(strings.Join(out, "\n"))
}
