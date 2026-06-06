package main

import "fmt"

// Observer receives push-based notifications from an Observable.
type Observer struct {
	next     func(int)
	complete func()
}

// Observable is a function that wires a producer to an observer on subscribe.
type Observable func(Observer)

// fromSlice emits each element synchronously, then completes.
func fromSlice(xs []int) Observable {
	return func(o Observer) {
		for _, x := range xs {
			o.next(x)
		}
		if o.complete != nil {
			o.complete()
		}
	}
}

// zip pairs values by index: it buffers each source in a per-source queue and,
// whenever both queues are non-empty, dequeues one from each and emits combine.
func zip(a, b Observable, combine func(int, int) int) Observable {
	return func(o Observer) {
		var qa, qb []int
		drain := func() {
			for len(qa) > 0 && len(qb) > 0 {
				x, y := qa[0], qb[0]
				qa, qb = qa[1:], qb[1:]
				o.next(combine(x, y))
			}
		}
		a(Observer{next: func(x int) { qa = append(qa, x); drain() }})
		b(Observer{next: func(y int) { qb = append(qb, y); drain() }})
	}
}

func main() {
	a := fromSlice([]int{1, 2, 3})
	b := fromSlice([]int{10, 20, 30})
	zipped := zip(a, b, func(x, y int) int { return x + y })
	zipped(Observer{next: func(v int) { fmt.Println(v) }})
}
