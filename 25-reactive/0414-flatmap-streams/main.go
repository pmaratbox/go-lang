package main

import (
	"container/heap"
	"fmt"
)

// task is a scheduled unit of work in virtual time.
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

// Scheduler is a virtual-time priority-queue scheduler.
type Scheduler struct {
	queue taskQueue
	clock int
	seq   int
}

func NewScheduler() *Scheduler {
	s := &Scheduler{}
	heap.Init(&s.queue)
	return s
}

func (s *Scheduler) Schedule(t int, cb func()) *task {
	tk := &task{time: t, seq: s.seq, cb: cb}
	s.seq++
	heap.Push(&s.queue, tk)
	return tk
}

func (s *Scheduler) Run() {
	for s.queue.Len() > 0 {
		tk := heap.Pop(&s.queue).(*task)
		if tk.dead {
			continue
		}
		s.clock = tk.time
		tk.cb()
	}
}

// Observer receives pushed values.
type Observer struct {
	next func(int)
}

// Observable is a push-based source: subscribe wires a producer to an observer.
type Observable struct {
	subscribe func(o *Observer)
}

// timed emits each value at its absolute virtual time.
func timed(s *Scheduler, events []struct {
	t int
	v int
}) *Observable {
	return &Observable{subscribe: func(o *Observer) {
		for _, e := range events {
			e := e
			s.Schedule(e.t, func() { o.next(e.v) })
		}
	}}
}

// flatMap maps each outer value to an inner observable and merges all inners
// concurrently (no cancellation).
func flatMap(source *Observable, project func(int) *Observable) *Observable {
	return &Observable{subscribe: func(o *Observer) {
		source.subscribe(&Observer{next: func(v int) {
			inner := project(v)
			inner.subscribe(&Observer{next: func(iv int) {
				o.next(iv)
			}})
		}})
	}}
}

func main() {
	s := NewScheduler()

	outer := timed(s, []struct {
		t int
		v int
	}{
		{10, 1},
		{20, 2},
	})

	merged := flatMap(outer, func(n int) *Observable {
		// inner schedules (now+5 -> n) and (now+30 -> n*10)
		now := s.clock
		return timed(s, []struct {
			t int
			v int
		}{
			{now + 5, n},
			{now + 30, n * 10},
		})
	})

	merged.subscribe(&Observer{next: func(v int) {
		fmt.Println(v)
	}})

	s.Run()
}
