package main

import (
	"fmt"
	"strings"
)

type node struct {
	val        int
	prev, next *node
}

func main() {
	n1 := &node{val: 1}
	n2 := &node{val: 2}
	n3 := &node{val: 3}
	n1.next = n2
	n2.prev = n1
	n2.next = n3
	n3.prev = n2

	var fwd []string
	for n := n1; n != nil; n = n.next {
		fwd = append(fwd, fmt.Sprint(n.val))
	}
	fmt.Println(strings.Join(fwd, " "))

	var bwd []string
	for n := n3; n != nil; n = n.prev {
		bwd = append(bwd, fmt.Sprint(n.val))
	}
	fmt.Println(strings.Join(bwd, " "))
}
