package main

import "fmt"

func main() {
	n := 3
	adj := make([][]int, n)
	add := func(u, v int) { adj[u] = append(adj[u], v) }
	add(0, 1)
	add(1, 2)
	add(2, 0)

	const (
		white = 0
		gray  = 1
		black = 2
	)
	color := make([]int, n)
	var dfs func(u int) bool
	dfs = func(u int) bool {
		color[u] = gray
		for _, v := range adj[u] {
			if color[v] == gray {
				return true
			}
			if color[v] == white && dfs(v) {
				return true
			}
		}
		color[u] = black
		return false
	}

	found := false
	for i := 0; i < n; i++ {
		if color[i] == white && dfs(i) {
			found = true
			break
		}
	}
	if found {
		fmt.Println("cycle")
	} else {
		fmt.Println("acyclic")
	}
}
