package main

import (
	"fmt"
	"strings"
)

func siftDown(a []int, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child > hi {
			return
		}
		if child+1 <= hi && a[child] < a[child+1] {
			child++
		}
		if a[root] >= a[child] {
			return
		}
		a[root], a[child] = a[child], a[root]
		root = child
	}
}

func heapSort(a []int) {
	n := len(a)
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(a, i, n-1)
	}
	for end := n - 1; end > 0; end-- {
		a[0], a[end] = a[end], a[0]
		siftDown(a, 0, end-1)
	}
}

func main() {
	a := []int{5, 3, 8, 1, 4}
	heapSort(a)

	out := make([]string, len(a))
	for i, v := range a {
		out[i] = fmt.Sprint(v)
	}
	fmt.Println(strings.Join(out, " "))
}
