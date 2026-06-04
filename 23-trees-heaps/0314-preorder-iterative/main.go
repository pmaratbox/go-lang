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
	stack := []*node{root}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out = append(out, fmt.Sprint(n.val))
		if n.right != nil {
			stack = append(stack, n.right)
		}
		if n.left != nil {
			stack = append(stack, n.left)
		}
	}
	fmt.Println(strings.Join(out, " "))
}
