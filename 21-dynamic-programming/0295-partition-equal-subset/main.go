package main

import "fmt"

func main() {
	nums := []int{1, 5, 11, 5}
	total := 0
	for _, n := range nums {
		total += n
	}
	if total%2 != 0 {
		fmt.Println("no")
		return
	}
	target := total / 2
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
