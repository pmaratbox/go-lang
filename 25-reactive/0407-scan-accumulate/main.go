package main

import "fmt"

// Observer receives pushed values, an optional error, and a completion signal.
type Observer struct {
	next     func(int)
	err      func(error)
	complete func()
}

// Observable is a producer: subscribe wires it to an observer.
type Observable func(Observer)

// of emits each value in order, then completes.
func of(values ...int) Observable {
	return func(o Observer) {
		for _, v := range values {
			o.next(v)
		}
		if o.complete != nil {
			o.complete()
		}
	}
}

// scan emits the running accumulation seeded with acc; for each value
// state = f(state, value) and the new state is emitted.
func scan(source Observable, acc int, f func(int, int) int) Observable {
	return func(o Observer) {
		state := acc
		source(Observer{
			next: func(v int) {
				state = f(state, v)
				o.next(state)
			},
			err:      o.err,
			complete: o.complete,
		})
	}
}

func main() {
	source := of(1, 2, 3, 4)
	running := scan(source, 0, func(state, v int) int { return state + v })
	running(Observer{
		next: func(v int) { fmt.Println(v) },
	})
}
