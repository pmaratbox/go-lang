package main

import "fmt"

func longestCommonSubstring(a, b string) string {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	best, end := 0, 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > best {
					best = dp[i][j]
					end = i
				}
			}
		}
	}
	return a[end-best : end]
}

func main() {
	fmt.Println(longestCommonSubstring("abcde", "xbcdy"))
}
