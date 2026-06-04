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

func search(n *node, v int) bool {
	for n != nil {
		if v == n.val {
			return true
		}
		if v < n.val {
			n = n.left
		} else {
			n = n.right
		}
	}
	return false
}

func main() {
	var root *node
	for _, v := range []int{5, 3, 8, 1, 4} {
		root = insert(root, v)
	}

	yesNo := func(ok bool) string {
		if ok {
			return "yes"
		}
		return "no"
	}

	fmt.Println(yesNo(search(root, 4)), yesNo(search(root, 6)))
}
