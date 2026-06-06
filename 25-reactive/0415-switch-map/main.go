package main

import (
	"container/heap"
	"fmt"
)

// task is a scheduled callback in virtual time.
type task struct {
	time int
	seq  int
	cb   func()
	dead bool
}

// taskQueue is a priority queue ordered by (time, seq).
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

// Scheduler is a virtual-time scheduler.
type Scheduler struct {
	now   int
	seq   int
	queue taskQueue
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

func (s *Scheduler) Cancel(tk *task) {
	if tk != nil {
		tk.dead = true
	}
}

func (s *Scheduler) Run() {
	for s.queue.Len() > 0 {
		tk := heap.Pop(&s.queue).(*task)
		if tk.dead {
			continue
		}
		s.now = tk.time
		tk.cb()
	}
}

// Observer receives pushed values.
type Observer struct {
	next func(int)
}

func main() {
	sched := NewScheduler()

	// outer: (10->1),(20->2)
	outer := func(obs Observer) {
		sched.Schedule(10, func() { obs.next(1) })
		sched.Schedule(20, func() { obs.next(2) })
	}

	// switchMap: cancel previous inner's pending emissions on each new outer value.
	// inner(n) schedules (now+5 -> n),(now+30 -> n*10).
	var pending []*task
	switched := func(obs Observer) {
		outer(Observer{next: func(v int) {
			// cancel previous inner subscription
			for _, tk := range pending {
				sched.Cancel(tk)
			}
			pending = pending[:0]
			// start new inner, recording its scheduled tasks
			start := sched.now
			pending = append(pending,
				sched.Schedule(start+5, func() { obs.next(v) }),
				sched.Schedule(start+30, func() { obs.next(v * 10) }),
			)
		}})
	}

	switched(Observer{next: func(v int) {
		fmt.Println(v)
	}})

	sched.Run()
}
