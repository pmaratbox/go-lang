package main

import "fmt"

type MultisetCount struct {
	counts map[int]int
}

func newMultisetCount() *MultisetCount {
	return &MultisetCount{counts: make(map[int]int)}
}

func (m *MultisetCount) add(v int) {
	m.counts[v]++
}

func (m *MultisetCount) remove(v int) {
	if m.counts[v] > 0 {
		m.counts[v]--
		if m.counts[v] == 0 {
			delete(m.counts, v)
		}
	}
}

func (m *MultisetCount) count(v int) int {
	return m.counts[v]
}

func main() {
	m := newMultisetCount()
	m.add(1)
	m.add(1)
	m.add(2)
	first := m.count(1)
	m.remove(1)
	second := m.count(1)
	fmt.Println(first, second)
}
