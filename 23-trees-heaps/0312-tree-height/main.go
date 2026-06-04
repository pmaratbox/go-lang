package main

import "fmt"

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

func height(n *node) int {
	if n == nil {
		return 0
	}
	return 1 + max(height(n.left), height(n.right))
}

func main() {
	var root *node
	for _, v := range []int{5, 3, 8, 1, 4} {
		root = insert(root, v)
	}
	fmt.Println(height(root))
}
