package main

import "fmt"

type node interface {
	accept(v *sumVisitor)
}

type leaf struct{ value int }

func (l leaf) accept(v *sumVisitor) { v.visitLeaf(l) }

type branch struct{ children []node }

func (b branch) accept(v *sumVisitor) {
	for _, c := range b.children {
		c.accept(v)
	}
}

type sumVisitor struct{ total int }

func (v *sumVisitor) visitLeaf(l leaf) { v.total += l.value }

func main() {
	tree := branch{children: []node{leaf{1}, leaf{2}, leaf{3}}}
	v := &sumVisitor{}
	tree.accept(v)
	fmt.Println(v.total)
}
