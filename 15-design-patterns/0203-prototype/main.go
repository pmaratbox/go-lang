package main

import "fmt"

type prototype struct{ value int }

func (p *prototype) clone() *prototype {
	c := *p
	return &c
}

func main() {
	orig := &prototype{value: 1}
	clone := orig.clone()
	clone.value = 2
	fmt.Println(orig.value, clone.value)
}
