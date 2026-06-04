package main

import (
	"fmt"
	"strings"
)

func rabinKarp(text, pat string) []int {
	const base, mod = 256, 1000000007
	n, m := len(text), len(pat)
	var res []int
	if m == 0 || n < m {
		return res
	}
	var patHash, winHash, pow int
	pow = 1
	for i := 0; i < m; i++ {
		patHash = (patHash*base + int(pat[i])) % mod
		winHash = (winHash*base + int(text[i])) % mod
		if i > 0 {
			pow = (pow * base) % mod
		}
	}
	for i := 0; i+m <= n; i++ {
		if winHash == patHash && text[i:i+m] == pat {
			res = append(res, i)
		}
		if i+m < n {
			winHash = ((winHash-int(text[i])*pow)%mod*base + int(text[i+m])) % mod
			winHash = (winHash%mod + mod) % mod
		}
	}
	return res
}

func main() {
	idx := rabinKarp("xabxab", "ab")
	parts := make([]string, len(idx))
	for i, v := range idx {
		parts[i] = fmt.Sprintf("%d", v)
	}
	fmt.Println(strings.Join(parts, " "))
}
