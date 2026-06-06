package main

import "fmt"

// Observer receives pushed values.
type Observer struct {
	next func(int)
}

// Subscription is returned by subscribe; calling unsubscribe stops delivery.
type Subscription struct {
	closed bool
}

func (s *Subscription) unsubscribe() { s.closed = true }

// Observable wires a producer to an observer when subscribed. The producer is
// given the live Subscription so it can check the closed flag before each next.
type Observable struct {
	produce func(o Observer, sub *Subscription)
}

// subscribe creates the Subscription first, builds the Observer from it (so the
// consumer can unsubscribe synchronously during delivery), then runs the
// producer and returns the Subscription.
func (ob Observable) subscribe(build func(sub *Subscription) Observer) *Subscription {
	sub := &Subscription{}
	o := build(sub)
	ob.produce(o, sub)
	return sub
}

func main() {
	source := Observable{produce: func(o Observer, sub *Subscription) {
		for _, v := range []int{1, 2, 3, 4} {
			if sub.closed {
				return // source checks the closed flag before each next
			}
			o.next(v)
		}
	}}

	count := 0
	source.subscribe(func(sub *Subscription) Observer {
		return Observer{next: func(v int) {
			fmt.Println(v)
			count++
			if count == 2 {
				sub.unsubscribe() // unsubscribe after receiving 2
			}
		}}
	})
}
