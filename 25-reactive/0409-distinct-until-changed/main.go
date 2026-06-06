package main

import "fmt"

// Observer receives pushed values from an Observable.
type Observer struct {
	next     func(int)
	complete func()
}

// Observable is a push-based source: subscribe wires a producer to an observer.
type Observable func(Observer)

// of emits the given values synchronously, then completes.
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

// distinctUntilChanged forwards a value only when it differs from the last emitted one.
func distinctUntilChanged(src Observable) Observable {
	return func(o Observer) {
		hasLast := false
		var last int
		src(Observer{
			next: func(v int) {
				if !hasLast || v != last {
					hasLast = true
					last = v
					o.next(v)
				}
			},
			complete: o.complete,
		})
	}
}

func main() {
	source := of(1, 1, 2, 2, 2, 3, 1)
	distinctUntilChanged(source)(Observer{
		next: func(v int) { fmt.Println(v) },
	})
}
