package main

import "fmt"

func boyerMoore(text, pat string) int {
	last := make(map[byte]int)
	for i := 0; i < len(pat); i++ {
		last[pat[i]] = i
	}
	n, m := len(text), len(pat)
	s := 0
	for s <= n-m {
		j := m - 1
		for j >= 0 && pat[j] == text[s+j] {
			j--
		}
		if j < 0 {
			return s
		}
		li, ok := last[text[s+j]]
		if !ok {
			li = -1
		}
		shift := j - li
		if shift < 1 {
			shift = 1
		}
		s += shift
	}
	return -1
}

func main() {
	fmt.Println(boyerMoore("zzabc", "abc"))
}
