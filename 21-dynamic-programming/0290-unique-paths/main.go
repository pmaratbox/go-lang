package main

import "fmt"

func main() {
	rows, cols := 3, 3
	dp := make([]int, cols)
	for i := range dp {
		dp[i] = 1
	}
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			dp[c] += dp[c-1]
		}
	}
	fmt.Println(dp[cols-1])
}
