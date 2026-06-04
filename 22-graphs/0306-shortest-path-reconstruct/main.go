package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type edge struct {
	to, w int
}

type item struct {
	node, dist int
}

type pq []item

func (p pq) Len() int           { return len(p) }
func (p pq) Less(i, j int) bool { return p[i].dist < p[j].dist }
func (p pq) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *pq) Push(x any)        { *p = append(*p, x.(item)) }
func (p *pq) Pop() any {
	old := *p
	n := len(old)
	it := old[n-1]
	*p = old[:n-1]
	return it
}

func main() {
	n := 4
	adj := make([][]edge, n)
	add := func(u, v, w int) { adj[u] = append(adj[u], edge{v, w}) }
	add(0, 1, 4)
	add(0, 2, 1)
	add(2, 1, 2)
	add(1, 3, 1)
	add(2, 3, 5)

	const inf = 1 << 30
	dist := make([]int, n)
	prev := make([]int, n)
	for i := range dist {
		dist[i] = inf
		prev[i] = -1
	}
	dist[0] = 0
	q := &pq{{0, 0}}
	for q.Len() > 0 {
		cur := heap.Pop(q).(item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range adj[cur.node] {
			if nd := cur.dist + e.w; nd < dist[e.to] {
				dist[e.to] = nd
				prev[e.to] = cur.node
				heap.Push(q, item{e.to, nd})
			}
		}
	}

	var path []int
	for v := 3; v != -1; v = prev[v] {
		path = append([]int{v}, path...)
	}
	parts := make([]string, len(path))
	for i, v := range path {
		parts[i] = fmt.Sprint(v)
	}
	fmt.Println(strings.Join(parts, " "))
}
