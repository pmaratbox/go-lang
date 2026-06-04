package main

import "fmt"

type parser struct {
	s   string
	pos int
}

func (p *parser) expr() int {
	v := p.term()
	for p.pos < len(p.s) && (p.s[p.pos] == '+' || p.s[p.pos] == '-') {
		op := p.s[p.pos]
		p.pos++
		t := p.term()
		if op == '+' {
			v += t
		} else {
			v -= t
		}
	}
	return v
}

func (p *parser) term() int {
	v := p.factor()
	for p.pos < len(p.s) && (p.s[p.pos] == '*' || p.s[p.pos] == '/') {
		op := p.s[p.pos]
		p.pos++
		f := p.factor()
		if op == '*' {
			v *= f
		} else {
			v /= f
		}
	}
	return v
}

func (p *parser) factor() int {
	n := 0
	for p.pos < len(p.s) && p.s[p.pos] >= '0' && p.s[p.pos] <= '9' {
		n = n*10 + int(p.s[p.pos]-'0')
		p.pos++
	}
	return n
}

func main() {
	p := &parser{s: "2+3*4"}
	fmt.Println(p.expr())
}
