package main

import (
	"fmt"
	"strings"
)

var adj = map[int][]int{
	0: {1, 2},
	1: {0, 3},
	2: {0, 3},
	3: {1, 2},
}

func dfs(cur int, visited map[int]bool, order *[]int) {
	visited[cur] = true
	*order = append(*order, cur)
	for _, n := range adj[cur] {
		if !visited[n] {
			dfs(n, visited, order)
		}
	}
}

func main() {
	visited := map[int]bool{}
	var order []int
	dfs(0, visited, &order)
	var parts []string
	for _, v := range order {
		parts = append(parts, fmt.Sprint(v))
	}
	fmt.Println(strings.Join(parts, " "))
}
