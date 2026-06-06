package main

import "fmt"

// Observer receives push-based notifications from an Observable.
type Observer struct {
	next     func(int)
	err      func(error)
	complete func()
}

// Observable is a function that wires a producer to an observer on subscribe.
type Observable func(Observer)

func (o Observable) subscribe(obs Observer) { o(obs) }

// catchError forwards next values, but on error subscribes to the fallback
// instead of propagating the error.
func catchError(source, fallback Observable) Observable {
	return func(obs Observer) {
		source.subscribe(Observer{
			next: obs.next,
			err: func(error) {
				fallback.subscribe(obs)
			},
			complete: obs.complete,
		})
	}
}

func main() {
	source := Observable(func(obs Observer) {
		obs.next(1)
		obs.next(2)
		obs.err(fmt.Errorf("boom"))
	})

	fallback := Observable(func(obs Observer) {
		obs.next(9)
		obs.complete()
	})

	catchError(source, fallback).subscribe(Observer{
		next:     func(v int) { fmt.Println(v) },
		err:      func(error) {},
		complete: func() {},
	})
}
