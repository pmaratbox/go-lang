package main

import "fmt"

type Expr interface{ eval() int }

type Num struct{ v int }

func (n Num) eval() int { return n.v }

type Add struct{ l, r Expr }

func (a Add) eval() int { return a.l.eval() + a.r.eval() }

func main() {
	e := Add{l: Num{1}, r: Num{2}}
	fmt.Println(e.eval())
}
