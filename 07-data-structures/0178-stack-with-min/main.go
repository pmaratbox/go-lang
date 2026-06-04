package main

import "fmt"

type minStack struct {
	data []int
	mins []int
}

func (s *minStack) push(v int) {
	s.data = append(s.data, v)
	if len(s.mins) == 0 || v < s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, v)
	} else {
		s.mins = append(s.mins, s.mins[len(s.mins)-1])
	}
}

func (s *minStack) getMin() int {
	return s.mins[len(s.mins)-1]
}

func main() {
	s := &minStack{}
	for _, v := range []int{3, 1, 2} {
		s.push(v)
	}
	fmt.Printf("min: %d\n", s.getMin())
}
