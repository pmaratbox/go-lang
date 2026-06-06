package main

import "fmt"

// Observer receives values pushed by an Observable.
type Observer struct {
	Next     func(int)
	Complete func()
}

// Observable is a function that, on subscribe, pushes values to an observer.
type Observable func(Observer)

func main() {
	// An Observable that emits 1, 2, 3 then completes.
	observableCreate := Observable(func(o Observer) {
		o.Next(1)
		o.Next(2)
		o.Next(3)
		o.Complete()
	})

	observableCreate(Observer{
		Next:     func(v int) { fmt.Println(v) },
		Complete: func() { fmt.Println("done") },
	})
}
