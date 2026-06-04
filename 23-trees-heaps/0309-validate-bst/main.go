package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func isBST(n *node, low, high int) bool {
	if n == nil {
		return true
	}
	if n.val <= low || n.val >= high {
		return false
	}
	return isBST(n.left, low, n.val) && isBST(n.right, n.val, high)
}

func main() {
	const minInt, maxInt = -1 << 62, 1 << 62

	// valid BST
	good := &node{5,
		&node{3, &node{1, nil, nil}, &node{4, nil, nil}},
		&node{8, nil, nil}}

	// out-of-place node: 6 sits in the left subtree of 5
	bad := &node{5,
		&node{3, &node{1, nil, nil}, &node{6, nil, nil}},
		&node{8, nil, nil}}

	yesNo := func(ok bool) string {
		if ok {
			return "yes"
		}
		return "no"
	}

	fmt.Println(yesNo(isBST(good, minInt, maxInt)), yesNo(isBST(bad, minInt, maxInt)))
}
