package main

import (
	"fmt"
	"strings"
)

type node struct {
	val         int
	left, right *node
}

func insert(n *node, v int) *node {
	if n == nil {
		return &node{val: v}
	}
	if v < n.val {
		n.left = insert(n.left, v)
	} else if v > n.val {
		n.right = insert(n.right, v)
	}
	return n
}

func main() {
	var root *node
	for _, v := range []int{5, 3, 8, 1, 4} {
		root = insert(root, v)
	}

	var out []string
	queue := []*node{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		out = append(out, fmt.Sprint(n.val))
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
	fmt.Println(strings.Join(out, " "))
}
