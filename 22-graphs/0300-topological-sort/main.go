package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	n := 4
	adj := make([][]int, n)
	indeg := make([]int, n)
	add := func(u, v int) {
		adj[u] = append(adj[u], v)
		indeg[v]++
	}
	add(0, 1)
	add(0, 2)
	add(1, 3)
	add(2, 3)

	var ready []int
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			ready = append(ready, i)
		}
	}
	var order []int
	for len(ready) > 0 {
		sort.Ints(ready)
		u := ready[0]
		ready = ready[1:]
		order = append(order, u)
		for _, v := range adj[u] {
			indeg[v]--
			if indeg[v] == 0 {
				ready = append(ready, v)
			}
		}
	}
	parts := make([]string, len(order))
	for i, v := range order {
		parts[i] = fmt.Sprint(v)
	}
	fmt.Println(strings.Join(parts, " "))
}
