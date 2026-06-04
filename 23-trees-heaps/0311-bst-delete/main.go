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

func minNode(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func delete(n *node, v int) *node {
	if n == nil {
		return nil
	}
	switch {
	case v < n.val:
		n.left = delete(n.left, v)
	case v > n.val:
		n.right = delete(n.right, v)
	default:
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		succ := minNode(n.right)
		n.val = succ.val
		n.right = delete(n.right, succ.val)
	}
	return n
}

func inorder(n *node, out *[]string) {
	if n == nil {
		return
	}
	inorder(n.left, out)
	*out = append(*out, fmt.Sprint(n.val))
	inorder(n.right, out)
}

func main() {
	var root *node
	for _, v := range []int{5, 3, 8, 1, 4} {
		root = insert(root, v)
	}
	root = delete(root, 3)

	var out []string
	inorder(root, &out)
	fmt.Println(strings.Join(out, " "))
}
