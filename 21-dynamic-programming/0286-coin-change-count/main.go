package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	target := 5
	dp := make([]int, target+1)
	dp[0] = 1
	for _, c := range coins {
		for i := c; i <= target; i++ {
			dp[i] += dp[i-c]
		}
	}
	fmt.Println(dp[target])
}
