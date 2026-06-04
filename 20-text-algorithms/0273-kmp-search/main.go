package main

import (
	"fmt"
	"strings"
)

func prefixFunction(s string) []int {
	pi := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		j := pi[i-1]
		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

func kmpSearch(text, pat string) []int {
	pi := prefixFunction(pat)
	var res []int
	j := 0
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != pat[j] {
			j = pi[j-1]
		}
		if text[i] == pat[j] {
			j++
		}
		if j == len(pat) {
			res = append(res, i-len(pat)+1)
			j = pi[j-1]
		}
	}
	return res
}

func main() {
	idx := kmpSearch("ababab", "ab")
	parts := make([]string, len(idx))
	for i, v := range idx {
		parts[i] = fmt.Sprintf("%d", v)
	}
	fmt.Println(strings.Join(parts, " "))
}
