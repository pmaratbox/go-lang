package main

import "fmt"

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for r := range dp {
		dp[r] = make([]int, cols)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			best := 0
			switch {
			case r == 0 && c == 0:
			case r == 0:
				best = dp[r][c-1]
			case c == 0:
				best = dp[r-1][c]
			default:
				best = dp[r-1][c]
				if dp[r][c-1] < best {
					best = dp[r][c-1]
				}
			}
			dp[r][c] = best + grid[r][c]
		}
	}
	fmt.Println(dp[rows-1][cols-1])
}
