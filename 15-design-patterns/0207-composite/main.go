package main

import "fmt"

type component interface {
	size() int
}

type leaf struct{ value int }

func (l leaf) size() int { return l.value }

type composite struct{ children []component }

func (c composite) size() int {
	total := 0
	for _, ch := range c.children {
		total += ch.size()
	}
	return total
}

func main() {
	tree := composite{children: []component{leaf{1}, leaf{2}, leaf{3}}}
	fmt.Println(tree.size())
}
