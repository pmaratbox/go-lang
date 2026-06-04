package main

import (
	"container/heap"
	"fmt"
)

type cell struct {
	r, c int
}

type item struct {
	pos  cell
	g, f int
}

type pq []item

func (p pq) Len() int           { return len(p) }
func (p pq) Less(i, j int) bool { return p[i].f < p[j].f }
func (p pq) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *pq) Push(x any)        { *p = append(*p, x.(item)) }
func (p *pq) Pop() any {
	old := *p
	n := len(old)
	it := old[n-1]
	*p = old[:n-1]
	return it
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	const size = 3
	start := cell{0, 0}
	goal := cell{2, 2}
	h := func(c cell) int { return abs(c.r-goal.r) + abs(c.c-goal.c) }

	best := map[cell]int{start: 0}
	q := &pq{{start, 0, h(start)}}
	dirs := []cell{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for q.Len() > 0 {
		cur := heap.Pop(q).(item)
		if cur.pos == goal {
			fmt.Println(cur.g)
			return
		}
		if cur.g > best[cur.pos] {
			continue
		}
		for _, d := range dirs {
			nr, nc := cur.pos.r+d.r, cur.pos.c+d.c
			if nr < 0 || nr >= size || nc < 0 || nc >= size {
				continue
			}
			np := cell{nr, nc}
			ng := cur.g + 1
			if prev, ok := best[np]; !ok || ng < prev {
				best[np] = ng
				heap.Push(q, item{np, ng, ng + h(np)})
			}
		}
	}
}
