package main

import (
	"fmt"
	"strings"
)

type minHeap struct {
	data []int
}

func (h *minHeap) push(v int) {
	h.data = append(h.data, v)
	i := len(h.data) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent] <= h.data[i] {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *minHeap) pop() int {
	n := len(h.data)
	top := h.data[0]
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]
	n--
	i := 0
	for {
		l, r, smallest := 2*i+1, 2*i+2, i
		if l < n && h.data[l] < h.data[smallest] {
			smallest = l
		}
		if r < n && h.data[r] < h.data[smallest] {
			smallest = r
		}
		if smallest == i {
			break
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
	return top
}

func main() {
	h := &minHeap{}
	for _, v := range []int{3, 1, 2} {
		h.push(v)
	}
	var parts []string
	for len(h.data) > 0 {
		parts = append(parts, fmt.Sprint(h.pop()))
	}
	fmt.Println(strings.Join(parts, " "))
}
