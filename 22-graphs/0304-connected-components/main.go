package main

import "fmt"

func main() {
	n := 5
	adj := make([][]int, n)
	add := func(u, v int) {
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	add(0, 1)
	add(1, 2)
	add(3, 4)

	seen := make([]bool, n)
	var dfs func(u int)
	dfs = func(u int) {
		seen[u] = true
		for _, v := range adj[u] {
			if !seen[v] {
				dfs(v)
			}
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		if !seen[i] {
			count++
			dfs(i)
		}
	}
	fmt.Println(count)
}
