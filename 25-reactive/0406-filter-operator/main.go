package main

import "fmt"

// Observer receives pushed values from an Observable.
type Observer struct {
	next     func(int)
	err      func(error)
	complete func()
}

// Observable is a push-based source: subscribing wires a producer to an Observer.
type Observable func(Observer)

// of emits the given values in order, then completes.
func of(values ...int) Observable {
	return func(o Observer) {
		for _, v := range values {
			o.next(v)
		}
		o.complete()
	}
}

// filter forwards a value only when pred(value) holds.
func filter(source Observable, pred func(int) bool) Observable {
	return func(o Observer) {
		source(Observer{
			next: func(v int) {
				if pred(v) {
					o.next(v)
				}
			},
			err:      o.err,
			complete: o.complete,
		})
	}
}

func main() {
	source := of(1, 2, 3, 4, 5, 6)
	even := filter(source, func(v int) bool { return v%2 == 0 })
	even(Observer{
		next:     func(v int) { fmt.Println(v) },
		err:      func(error) {},
		complete: func() {},
	})
}
