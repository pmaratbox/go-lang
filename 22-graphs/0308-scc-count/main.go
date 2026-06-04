package main

import "fmt"

func main() {
	n := 4
	adj := make([][]int, n)
	radj := make([][]int, n)
	add := func(u, v int) {
		adj[u] = append(adj[u], v)
		radj[v] = append(radj[v], u)
	}
	add(0, 1)
	add(1, 2)
	add(2, 0)
	add(2, 3)

	// Kosaraju: first pass to get finish order.
	visited := make([]bool, n)
	var order []int
	var dfs1 func(u int)
	dfs1 = func(u int) {
		visited[u] = true
		for _, v := range adj[u] {
			if !visited[v] {
				dfs1(v)
			}
		}
		order = append(order, u)
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs1(i)
		}
	}

	// Second pass on reverse graph in reverse finish order.
	comp := make([]bool, n)
	var dfs2 func(u int)
	dfs2 = func(u int) {
		comp[u] = true
		for _, v := range radj[u] {
			if !comp[v] {
				dfs2(v)
			}
		}
	}
	count := 0
	for i := len(order) - 1; i >= 0; i-- {
		u := order[i]
		if !comp[u] {
			count++
			dfs2(u)
		}
	}
	fmt.Println(count)
}
