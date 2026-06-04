package main

import "fmt"

type edge struct {
	to, w int
}

func main() {
	n := 4
	adj := make([][]edge, n)
	add := func(u, v, w int) {
		adj[u] = append(adj[u], edge{v, w})
		adj[v] = append(adj[v], edge{u, w})
	}
	add(0, 1, 1)
	add(1, 2, 2)
	add(2, 3, 3)

	const inf = 1 << 30
	key := make([]int, n)
	inMST := make([]bool, n)
	for i := range key {
		key[i] = inf
	}
	key[0] = 0
	total := 0
	for i := 0; i < n; i++ {
		u := -1
		for v := 0; v < n; v++ {
			if !inMST[v] && (u == -1 || key[v] < key[u]) {
				u = v
			}
		}
		inMST[u] = true
		total += key[u]
		for _, e := range adj[u] {
			if !inMST[e.to] && e.w < key[e.to] {
				key[e.to] = e.w
			}
		}
	}
	fmt.Println(total)
}
