package main

import "fmt"

type Node interface {
	Eval() int
}

type Num struct{ v int }

func (n Num) Eval() int { return n.v }

type Add struct{ l, r Node }

func (a Add) Eval() int { return a.l.Eval() + a.r.Eval() }

type Mul struct{ l, r Node }

func (m Mul) Eval() int { return m.l.Eval() * m.r.Eval() }

func main() {
	// 1 + 2 * 3
	ast := Add{Num{1}, Mul{Num{2}, Num{3}}}
	fmt.Println(ast.Eval())
}
