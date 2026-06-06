package main

import (
	"container/heap"
	"fmt"
)

// task is a scheduled callback at a virtual time.
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

// scheduler is a virtual-time priority queue of callbacks.
type scheduler struct {
	queue taskQueue
	now   int
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

func (s *scheduler) cancel(tk *task) {
	if tk != nil {
		tk.dead = true
	}
}

func (s *scheduler) run() {
	for s.queue.Len() > 0 {
		tk := heap.Pop(&s.queue).(*task)
		if tk.dead {
			continue
		}
		s.now = tk.time
		tk.cb()
	}
}

// observer receives pushed values and a completion signal.
type observer struct {
	next     func(string)
	complete func()
}

// debounce emits a value only after a quiet gap of window ticks.
func debounce(s *scheduler, window int, out observer) observer {
	var pending *task
	var hasValue bool
	var lastValue string
	return observer{
		next: func(v string) {
			s.cancel(pending)
			lastValue = v
			hasValue = true
			pending = s.schedule(s.now+window, func() {
				hasValue = false
				out.next(lastValue)
			})
		},
		complete: func() {
			_ = hasValue
			out.complete()
		},
	}
}

func main() {
	s := newScheduler()

	sink := debounce(s, 30, observer{
		next:     func(v string) { fmt.Println(v) },
		complete: func() {},
	})

	// Source schedules ("a"@10),("b"@20),("c"@100).
	type ev struct {
		t int
		v string
	}
	events := []ev{{10, "a"}, {20, "b"}, {100, "c"}}
	for _, e := range events {
		e := e
		s.schedule(e.t, func() { sink.next(e.v) })
	}
	s.schedule(100, func() { sink.complete() })

	s.run()
}
