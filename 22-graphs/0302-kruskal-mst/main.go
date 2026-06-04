package main

import (
	"fmt"
	"sort"
)

type edge struct {
	u, v, w int
}

type dsu struct {
	parent []int
}

func newDSU(n int) *dsu {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &dsu{p}
}

func (d *dsu) find(x int) int {
	for d.parent[x] != x {
		d.parent[x] = d.parent[d.parent[x]]
		x = d.parent[x]
	}
	return x
}

func (d *dsu) union(a, b int) bool {
	ra, rb := d.find(a), d.find(b)
	if ra == rb {
		return false
	}
	d.parent[ra] = rb
	return true
}

func main() {
	n := 3
	edges := []edge{{0, 1, 1}, {1, 2, 2}, {0, 2, 3}}
	sort.Slice(edges, func(i, j int) bool { return edges[i].w < edges[j].w })
	d := newDSU(n)
	total := 0
	for _, e := range edges {
		if d.union(e.u, e.v) {
			total += e.w
		}
	}
	fmt.Println(total)
}
