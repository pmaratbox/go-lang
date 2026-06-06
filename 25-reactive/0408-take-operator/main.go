package main

import "fmt"

// Observer receives push-based notifications from a source.
type Observer struct {
	next     func(int)
	complete func()
}

// Subscription lets a downstream operator stop an upstream source.
type Subscription struct {
	unsubscribe func()
}

// naturals is an unbounded source emitting 1,2,3,... It keeps emitting
// only while the subscription is active, so take can stop it.
func naturals(obs Observer) Subscription {
	stopped := false
	sub := Subscription{unsubscribe: func() { stopped = true }}
	// Drive synchronously: emit until unsubscribed. A downstream
	// operator (take) flips stopped during obs.next, ending the loop.
	for i := 1; !stopped; i++ {
		obs.next(i)
	}
	return sub
}

// take(source, n) forwards the first n emissions, then completes and
// unsubscribes the source so the infinite producer stops.
func take(source func(Observer) Subscription, n int) func(Observer) Subscription {
	return func(downstream Observer) Subscription {
		count := 0
		var sub Subscription
		inner := Observer{
			next: func(v int) {
				if count < n {
					downstream.next(v)
					count++
					if count == n {
						downstream.complete()
						if sub.unsubscribe != nil {
							sub.unsubscribe()
						}
					}
				}
			},
			complete: downstream.complete,
		}
		sub = source(inner)
		return sub
	}
}

func main() {
	taken := take(naturals, 3)
	taken(Observer{
		next:     func(v int) { fmt.Println(v) },
		complete: func() { fmt.Println("completed") },
	})
}
