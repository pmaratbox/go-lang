package main

import (
	"fmt"
	"strings"
)

type rangeIterator struct {
	current, end int
}

func (it *rangeIterator) hasNext() bool { return it.current <= it.end }

func (it *rangeIterator) next() int {
	v := it.current
	it.current++
	return v
}

func main() {
	it := &rangeIterator{current: 1, end: 3}
	var parts []string
	for it.hasNext() {
		parts = append(parts, fmt.Sprint(it.next()))
	}
	fmt.Println(strings.Join(parts, " "))
}
