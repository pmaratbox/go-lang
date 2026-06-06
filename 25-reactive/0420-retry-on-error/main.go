package main

import "fmt"

// Observer receives push-based notifications.
type Observer struct {
	next     func(v int)
	err      func(e error)
	complete func()
}

// Observable wires a producer to an observer when subscribed.
type Observable struct {
	subscribe func(o Observer)
}

// retry resubscribes to the source on error up to n times.
func retry(source Observable, n int) Observable {
	return Observable{subscribe: func(o Observer) {
		var attempt func(remaining int)
		attempt = func(remaining int) {
			source.subscribe(Observer{
				next:     o.next,
				complete: o.complete,
				err: func(e error) {
					if remaining > 0 {
						attempt(remaining - 1)
					} else {
						o.err(e)
					}
				},
			})
		}
		attempt(n)
	}}
}

func main() {
	count := 0
	source := Observable{subscribe: func(o Observer) {
		count++
		k := count
		fmt.Printf("attempt %d\n", k)
		if k < 3 {
			o.err(fmt.Errorf("fail %d", k))
		} else {
			fmt.Println("ok")
			o.complete()
		}
	}}

	retry(source, 2).subscribe(Observer{
		next:     func(v int) {},
		err:      func(e error) {},
		complete: func() {},
	})
}
