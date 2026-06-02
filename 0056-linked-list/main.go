package main

import (
	"fmt"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

func main() {
	head := &Node{1, &Node{2, &Node{3, nil}}}

	parts := []string{}
	for node := head; node != nil; node = node.next {
		parts = append(parts, fmt.Sprint(node.value))
	}
	fmt.Println(strings.Join(parts, " -> "))
}
