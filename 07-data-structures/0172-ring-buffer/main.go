package main

import (
	"fmt"
	"strings"
)

type ring struct {
	buf   []int
	head  int
	count int
}

func newRing(cap int) *ring {
	return &ring{buf: make([]int, cap)}
}

func (r *ring) push(v int) {
	idx := (r.head + r.count) % len(r.buf)
	if r.count == len(r.buf) {
		r.buf[r.head] = v
		r.head = (r.head + 1) % len(r.buf)
	} else {
		r.buf[idx] = v
		r.count++
	}
}

func (r *ring) contents() []int {
	out := make([]int, r.count)
	for i := 0; i < r.count; i++ {
		out[i] = r.buf[(r.head+i)%len(r.buf)]
	}
	return out
}

func main() {
	r := newRing(3)
	for _, v := range []int{1, 2, 3, 4, 5} {
		r.push(v)
	}
	var parts []string
	for _, v := range r.contents() {
		parts = append(parts, fmt.Sprint(v))
	}
	fmt.Println(strings.Join(parts, " "))
}
