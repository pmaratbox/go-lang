package main

import "fmt"

func main() {
	nums := []int{3, 34, 4, 12, 5, 2}
	target := 9
	dp := make([]bool, target+1)
	dp[0] = true
	for _, n := range nums {
		for s := target; s >= n; s-- {
			if dp[s-n] {
				dp[s] = true
			}
		}
	}
	if dp[target] {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
