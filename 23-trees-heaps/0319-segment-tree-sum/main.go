package main

import "fmt"

type segTree struct {
	n    int
	tree []int
}

func build(data []int) *segTree {
	n := len(data)
	st := &segTree{n: n, tree: make([]int, 2*n)}
	for i, v := range data {
		st.tree[n+i] = v
	}
	for i := n - 1; i > 0; i-- {
		st.tree[i] = st.tree[2*i] + st.tree[2*i+1]
	}
	return st
}

// query sum over [l, r] inclusive
func (st *segTree) query(l, r int) int {
	l += st.n
	r += st.n + 1
	sum := 0
	for l < r {
		if l&1 == 1 {
			sum += st.tree[l]
			l++
		}
		if r&1 == 1 {
			r--
			sum += st.tree[r]
		}
		l >>= 1
		r >>= 1
	}
	return sum
}

func main() {
	st := build([]int{1, 2, 3, 4, 5})
	fmt.Println(st.query(1, 3))
}
