package main

import "fmt"

// Observer receives push-based notifications from an Observable.
type Observer struct {
	next     func(int)
	complete func()
}

// Observable is a function that wires a producer to an observer when subscribed.
type Observable func(Observer)

// fromValues builds an Observable that synchronously emits the given values
// and then completes.
func fromValues(values ...int) Observable {
	return func(o Observer) {
		for _, v := range values {
			o.next(v)
		}
		o.complete()
	}
}

// concat subscribes to a, and only after a completes subscribes to b;
// it completes after b completes.
func concat(a, b Observable) Observable {
	return func(o Observer) {
		a(Observer{
			next: o.next,
			complete: func() {
				b(Observer{
					next:     o.next,
					complete: o.complete,
				})
			},
		})
	}
}

func main() {
	a := fromValues(1, 2)
	b := fromValues(3, 4)

	concat(a, b)(Observer{
		next:     func(v int) { fmt.Println(v) },
		complete: func() {},
	})
}
