package main

import (
	"container/heap"
	"fmt"
)

// task is a unit of work scheduled at a virtual time.
type task struct {
	time      int
	seq       int
	cb        func()
	cancelled bool
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

// scheduler is a virtual-time scheduler.
type scheduler struct {
	now   int
	seq   int
	queue taskQueue
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

func (s *scheduler) cancel(tk *task) { tk.cancelled = true }

func (s *scheduler) run() {
	for s.queue.Len() > 0 {
		tk := heap.Pop(&s.queue).(*task)
		if tk.cancelled {
			continue
		}
		s.now = tk.time
		tk.cb()
	}
}

// observer receives pushed values.
type observer struct {
	next func(string)
}

// throttle applies leading-edge throttling: emit a value, then suppress
// further values for window ticks.
func throttle(s *scheduler, window int, out *observer) *observer {
	blockUntil := 0
	return &observer{
		next: func(v string) {
			if s.now >= blockUntil {
				out.next(v)
				blockUntil = s.now + window
			}
		},
	}
}

func main() {
	s := newScheduler()

	sink := &observer{next: func(v string) { fmt.Println(v) }}
	thr := throttle(s, 30, sink)

	// Source schedules its emissions at virtual times.
	events := []struct {
		time int
		val  string
	}{
		{10, "a"},
		{20, "b"},
		{100, "c"},
		{110, "d"},
	}
	for _, e := range events {
		val := e.val
		s.schedule(e.time, func() { thr.next(val) })
	}

	s.run()
}
