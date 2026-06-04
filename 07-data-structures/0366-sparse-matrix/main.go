package main

import "fmt"

type cell struct {
	row, col int
}

type SparseMatrix struct {
	values map[cell]int
}

func newSparseMatrix() *SparseMatrix {
	return &SparseMatrix{values: make(map[cell]int)}
}

func (s *SparseMatrix) set(row, col, v int) {
	if v == 0 {
		delete(s.values, cell{row, col})
		return
	}
	s.values[cell{row, col}] = v
}

func (s *SparseMatrix) get(row, col int) int {
	return s.values[cell{row, col}]
}

func main() {
	s := newSparseMatrix()
	s.set(1, 1, 5)
	fmt.Println(s.get(1, 1), s.get(0, 0))
}
