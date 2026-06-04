package main

import (
	"fmt"
	"strings"
)

type Deque struct {
	data []int
}

func (d *Deque) pushBack(v int) {
	d.data = append(d.data, v)
}

func (d *Deque) pushFront(v int) {
	d.data = append([]int{v}, d.data...)
}

func main() {
	d := &Deque{}
	d.pushBack(1)
	d.pushBack(2)
	d.pushFront(0)
	parts := make([]string, 0, len(d.data))
	for _, v := range d.data {
		parts = append(parts, fmt.Sprint(v))
	}
	fmt.Println(strings.Join(parts, " "))
}
