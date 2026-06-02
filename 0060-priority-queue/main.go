package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	h := &IntHeap{}
	heap.Init(h)
	for _, n := range []int{3, 1, 2} {
		heap.Push(h, n)
	}

	out := []string{}
	for h.Len() > 0 {
		out = append(out, fmt.Sprint(heap.Pop(h)))
	}
	fmt.Println(strings.Join(out, " "))
}
