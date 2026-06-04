package main

import "fmt"

type fenwick struct {
	tree []int
}

func newFenwick(n int) *fenwick {
	return &fenwick{tree: make([]int, n+1)}
}

// add v at 1-based index i
func (f *fenwick) update(i, v int) {
	for ; i < len(f.tree); i += i & (-i) {
		f.tree[i] += v
	}
}

// prefix sum of the first i elements (1-based, inclusive)
func (f *fenwick) prefix(i int) int {
	sum := 0
	for ; i > 0; i -= i & (-i) {
		sum += f.tree[i]
	}
	return sum
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	f := newFenwick(len(data))
	for i, v := range data {
		f.update(i+1, v)
	}
	fmt.Println(f.prefix(4))
}
