package main

import "fmt"

// Observer receives values pushed by an Observable.
type Observer struct {
	next     func(int)
	err      func(error)
	complete func()
}

// Observable is a push-based source: subscribe wires a producer to an observer.
type Observable struct {
	subscribe func(Observer)
}

// fromSlice builds an Observable that synchronously emits each value, then completes.
func fromSlice(values []int) Observable {
	return Observable{subscribe: func(o Observer) {
		for _, v := range values {
			o.next(v)
		}
		if o.complete != nil {
			o.complete()
		}
	}}
}

// mapOp returns a new Observable whose next forwards f(value).
func mapOp(source Observable, f func(int) int) Observable {
	return Observable{subscribe: func(o Observer) {
		source.subscribe(Observer{
			next:     func(v int) { o.next(f(v)) },
			err:      o.err,
			complete: o.complete,
		})
	}}
}

func main() {
	source := fromSlice([]int{1, 2, 3, 4})
	doubled := mapOp(source, func(x int) int { return x * 2 })
	doubled.subscribe(Observer{
		next: func(v int) { fmt.Println(v) },
	})
}
