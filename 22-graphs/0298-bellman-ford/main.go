package main

import (
	"fmt"
	"strings"
)

type edge struct {
	from, to, w int
}

func bellmanFord(edges []edge, src, n int) []int {
	const inf = 1 << 30
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[src] = 0
	for i := 0; i < n-1; i++ {
		for _, e := range edges {
			if dist[e.from] != inf && dist[e.from]+e.w < dist[e.to] {
				dist[e.to] = dist[e.from] + e.w
			}
		}
	}
	return dist
}

func main() {
	n := 3
	edges := []edge{{0, 1, 1}, {1, 2, -2}, {0, 2, 4}}
	dist := bellmanFord(edges, 0, n)
	parts := make([]string, n)
	for i, d := range dist {
		parts[i] = fmt.Sprint(d)
	}
	fmt.Println(strings.Join(parts, " "))
}
