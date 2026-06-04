package main

import (
	"fmt"
	"strings"
)

type CircularQueue struct {
	data        []int
	head, count int
}

func newCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{data: make([]int, capacity)}
}

func (q *CircularQueue) enqueue(v int) bool {
	if q.count == len(q.data) {
		return false
	}
	tail := (q.head + q.count) % len(q.data)
	q.data[tail] = v
	q.count++
	return true
}

func (q *CircularQueue) dequeue() (int, bool) {
	if q.count == 0 {
		return 0, false
	}
	v := q.data[q.head]
	q.head = (q.head + 1) % len(q.data)
	q.count--
	return v, true
}

func (q *CircularQueue) values() []int {
	out := make([]int, q.count)
	for i := 0; i < q.count; i++ {
		out[i] = q.data[(q.head+i)%len(q.data)]
	}
	return out
}

func main() {
	q := newCircularQueue(3)
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.dequeue()
	q.enqueue(4)
	parts := make([]string, 0)
	for _, v := range q.values() {
		parts = append(parts, fmt.Sprint(v))
	}
	fmt.Println(strings.Join(parts, " "))
}
