package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	target := 11
	const inf = 1 << 30
	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = inf
		for _, c := range coins {
			if c <= i && dp[i-c]+1 < dp[i] {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	fmt.Println(dp[target])
}
