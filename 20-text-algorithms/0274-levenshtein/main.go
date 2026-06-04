package main

import "fmt"

func levenshtein(a, b string) int {
	m, n := len(a), len(b)
	prev := make([]int, n+1)
	for j := 0; j <= n; j++ {
		prev[j] = j
	}
	for i := 1; i <= m; i++ {
		cur := make([]int, n+1)
		cur[0] = i
		for j := 1; j <= n; j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}
			cur[j] = min(prev[j]+1, min(cur[j-1]+1, prev[j-1]+cost))
		}
		prev = cur
	}
	return prev[n]
}

func main() {
	fmt.Println(levenshtein("kitten", "sitting"))
}
