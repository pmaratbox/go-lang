package main

import "fmt"

type trieNode struct {
	children map[rune]*trieNode
	end      bool
}

func newNode() *trieNode {
	return &trieNode{children: make(map[rune]*trieNode)}
}

func (t *trieNode) insert(word string) {
	cur := t
	for _, r := range word {
		if cur.children[r] == nil {
			cur.children[r] = newNode()
		}
		cur = cur.children[r]
	}
	cur.end = true
}

func (t *trieNode) search(word string) bool {
	cur := t
	for _, r := range word {
		if cur.children[r] == nil {
			return false
		}
		cur = cur.children[r]
	}
	return cur.end
}

func yesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func main() {
	root := newNode()
	root.insert("cat")
	root.insert("car")
	fmt.Println(yesNo(root.search("car")), yesNo(root.search("can")))
}
