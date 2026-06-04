package main

import (
	"fmt"
	"sort"
	"strings"
)

func suffixArray(s string) []int {
	sa := make([]int, len(s))
	for i := range sa {
		sa[i] = i
	}
	sort.Slice(sa, func(i, j int) bool {
		return s[sa[i]:] < s[sa[j]:]
	})
	return sa
}

func main() {
	sa := suffixArray("banana")
	parts := make([]string, len(sa))
	for i, v := range sa {
		parts[i] = fmt.Sprintf("%d", v)
	}
	fmt.Println(strings.Join(parts, " "))
}
