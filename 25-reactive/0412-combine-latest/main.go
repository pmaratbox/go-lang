package main

import (
	"container/heap"
	"fmt"
)

// task is one scheduled callback at a virtual time.
type task struct {
	time int
	seq  int
	run  func()
	dead bool
}

// taskQueue is a min-heap ordered by (time, seq).
type taskQueue []*task

func (q taskQueue) Len() int { return len(q) }
func (q taskQueue) Less(i, j int) bool {
	if q[i].time != q[j].time {
		return q[i].time < q[j].time
	}
	return q[i].seq < q[j].seq
}
func (q taskQueue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q *taskQueue) Push(x any)   { *q = append(*q, x.(*task)) }
func (q *taskQueue) Pop() any {
	old := *q
	n := len(old)
	t := old[n-1]
	*q = old[:n-1]
	return t
}

// scheduler is a virtual-time priority queue of callbacks.
type scheduler struct {
	queue taskQueue
	clock int
	seq   int
}

func newScheduler() *scheduler {
	s := &scheduler{}
	heap.Init(&s.queue)
	return s
}

func (s *scheduler) schedule(t int, cb func()) *task {
	tk := &task{time: t, seq: s.seq, run: cb}
	s.seq++
	heap.Push(&s.queue, tk)
	return tk
}

func (s *scheduler) run() {
	for s.queue.Len() > 0 {
		tk := heap.Pop(&s.queue).(*task)
		if tk.dead {
			continue
		}
		s.clock = tk.time
		tk.run()
	}
}

// observable is a push-based stream: subscribe wires a producer to an observer.
type observable func(next func(int))

// timedSource emits its (time, value) events at the given virtual times.
func timedSource(s *scheduler, events [][2]int) observable {
	return func(next func(int)) {
		for _, ev := range events {
			t, v := ev[0], ev[1]
			s.schedule(t, func() { next(v) })
		}
	}
}

// combineLatest emits the pair of latest values whenever either source
// emits, once both sources have emitted at least once.
func combineLatest(a, b observable, emit func(int, int)) {
	var latestA, latestB int
	var hasA, hasB bool
	tryEmit := func() {
		if hasA && hasB {
			emit(latestA, latestB)
		}
	}
	a(func(v int) { latestA, hasA = v, true; tryEmit() })
	b(func(v int) { latestB, hasB = v, true; tryEmit() })
}

func main() {
	s := newScheduler()

	a := timedSource(s, [][2]int{{1, 1}, {3, 2}})
	b := timedSource(s, [][2]int{{2, 10}})

	combineLatest(a, b, func(x, y int) {
		fmt.Printf("(%d, %d)\n", x, y)
	})

	s.run()
}
