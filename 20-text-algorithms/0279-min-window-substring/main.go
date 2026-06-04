package main

import "fmt"

func minWindow(s, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	missing := len(t)
	start, end := 0, len(s)+1
	left := 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		if need[c] > 0 {
			missing--
		}
		need[c]--
		for missing == 0 {
			if right-left+1 < end-start {
				start, end = left, right+1
			}
			need[s[left]]++
			if need[s[left]] > 0 {
				missing++
			}
			left++
		}
	}
	if end == len(s)+1 {
		return ""
	}
	return s[start:end]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
