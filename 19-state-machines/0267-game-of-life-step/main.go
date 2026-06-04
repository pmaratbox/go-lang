package main

import "fmt"

func main() {
	const n = 3
	grid := [n][n]bool{
		{false, true, false},
		{false, true, false},
		{false, true, false},
	}
	next := [n][n]bool{}
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			live := 0
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					}
					nr, nc := r+dr, c+dc
					if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] {
						live++
					}
				}
			}
			if grid[r][c] {
				next[r][c] = live == 2 || live == 3
			} else {
				next[r][c] = live == 3
			}
		}
	}
	for r := 0; r < n; r++ {
		line := make([]byte, n)
		for c := 0; c < n; c++ {
			if next[r][c] {
				line[c] = '#'
			} else {
				line[c] = '.'
			}
		}
		fmt.Println(string(line))
	}
}
