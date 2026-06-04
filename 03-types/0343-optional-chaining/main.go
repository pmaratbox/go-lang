package main

import "fmt"

type C struct{ v int }
type B struct{ c *C }
type A struct{ b *B }

func read(a *A) int {
	if a != nil && a.b != nil && a.b.c != nil {
		return a.b.c.v
	}
	return 0
}

func main() {
	present := &A{b: &B{c: &C{v: 5}}}
	absent := &A{b: &B{}}
	fmt.Println(read(present), read(absent))
}
