package main

import "fmt"

type node struct {
	children []*node
}

var best int

// returns height in edges, updating best with the diameter
func depth(n *node) int {
	if n == nil {
		return 0
	}
	var top, second int
	for _, c := range n.children {
		h := depth(c)
		if h > top {
			second = top
			top = h
		} else if h > second {
			second = h
		}
	}
	if top+second > best {
		best = top + second
	}
	return top + 1
}

func main() {
	c := &node{}
	d := &node{}
	a := &node{children: []*node{c, d}}
	b := &node{}
	root := &node{children: []*node{a, b}}

	depth(root)
	fmt.Println(best)
}
