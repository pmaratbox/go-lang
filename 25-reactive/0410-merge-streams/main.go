package main

import (
	"container/heap"
	"fmt"
)

// task is a unit of scheduled work in virtual time.
type task struct {
	time int
	seq  int
	cb   func()
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
func (q taskQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *taskQueue) Push(x any)        { *q = append(*q, x.(*task)) }
func (q *taskQueue) Pop() any {
	old := *q
	n := len(old)
	t := old[n-1]
	old[n-1] = nil
	*q = old[:n-1]
	return t
}

// scheduler runs scheduled callbacks in virtual-time order.
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
	tk := &task{time: t, seq: s.seq, cb: cb}
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
		tk.cb()
	}
}

// observer receives pushed values.
type observer struct {
	next func(int)
}

// observable wires a producer to an observer when subscribed.
type observable struct {
	subscribe func(obs observer)
}

// timed builds an observable that emits values at given virtual times.
func timed(s *scheduler, events [][2]int) observable {
	return observable{subscribe: func(obs observer) {
		for _, e := range events {
			t, v := e[0], e[1]
			s.schedule(t, func() { obs.next(v) })
		}
	}}
}

// merge subscribes to both sources onto the same observer.
func merge(a, b observable) observable {
	return observable{subscribe: func(obs observer) {
		a.subscribe(obs)
		b.subscribe(obs)
	}}
}

func main() {
	s := newScheduler()

	a := timed(s, [][2]int{{10, 1}, {30, 3}, {50, 5}})
	b := timed(s, [][2]int{{20, 2}, {40, 4}, {60, 6}})

	merge(a, b).subscribe(observer{next: func(v int) {
		fmt.Println(v)
	}})

	s.run()
}
