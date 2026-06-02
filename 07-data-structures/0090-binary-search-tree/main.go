package main

import (
	"fmt"
	"strings"
)

type Node struct {
	value       int
	left, right *Node
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{value: value}
	}
	if value < root.value {
		root.left = insert(root.left, value)
	} else {
		root.right = insert(root.right, value)
	}
	return root
}

func inorder(root *Node, out *[]string) {
	if root == nil {
		return
	}
	inorder(root.left, out)
	*out = append(*out, fmt.Sprint(root.value))
	inorder(root.right, out)
}

func main() {
	var root *Node
	for _, v := range []int{5, 3, 8, 1, 4} {
		root = insert(root, v)
	}
	out := []string{}
	inorder(root, &out)
	fmt.Println(strings.Join(out, " "))
}
