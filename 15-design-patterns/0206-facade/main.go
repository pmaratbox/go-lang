package main

import "fmt"

type sub1 struct{}
type sub2 struct{}
type sub3 struct{}

func (sub1) init() {}
func (sub2) init() {}
func (sub3) init() {}

type facade struct {
	a sub1
	b sub2
	c sub3
}

func (f facade) start() string {
	f.a.init()
	f.b.init()
	f.c.init()
	return "ready"
}

func main() {
	f := facade{}
	fmt.Println(f.start())
}
