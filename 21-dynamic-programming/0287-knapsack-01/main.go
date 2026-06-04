package main

import "fmt"

func main() {
	type item struct{ w, v int }
	items := []item{{2, 3}, {3, 4}, {4, 5}}
	capacity := 5
	dp := make([]int, capacity+1)
	for _, it := range items {
		for c := capacity; c >= it.w; c-- {
			if dp[c-it.w]+it.v > dp[c] {
				dp[c] = dp[c-it.w] + it.v
			}
		}
	}
	fmt.Println(dp[capacity])
}
