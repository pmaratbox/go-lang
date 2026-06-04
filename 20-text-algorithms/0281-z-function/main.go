package main

import (
	"fmt"
	"strings"
)

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i < r {
			z[i] = min(r-i, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i] > r {
			l, r = i, i+z[i]
		}
	}
	return z
}

func main() {
	z := zFunction("aaaa")
	parts := make([]string, 0, len(z)-1)
	for _, v := range z[1:] {
		parts = append(parts, fmt.Sprintf("%d", v))
	}
	fmt.Println(strings.Join(parts, " "))
}
