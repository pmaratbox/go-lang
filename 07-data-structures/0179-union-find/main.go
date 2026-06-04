package main

import "fmt"

type unionFind struct {
	parent []int
}

func newUF(n int) *unionFind {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &unionFind{parent: p}
}

func (u *unionFind) find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *unionFind) union(a, b int) {
	u.parent[u.find(a)] = u.find(b)
}

func (u *unionFind) connected(a, b int) bool {
	return u.find(a) == u.find(b)
}

func yesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func main() {
	uf := newUF(4)
	uf.union(0, 1)
	uf.union(2, 3)
	fmt.Println(yesNo(uf.connected(0, 1)), yesNo(uf.connected(0, 2)))
}
