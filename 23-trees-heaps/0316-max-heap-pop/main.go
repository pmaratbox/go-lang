package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)         { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func main() {
	h := &maxHeap{3, 1, 4, 1, 5}
	heap.Init(h)

	var out []string
	for i := 0; i < 3; i++ {
		out = append(out, fmt.Sprint(heap.Pop(h)))
	}
	fmt.Println(strings.Join(out, " "))
}
