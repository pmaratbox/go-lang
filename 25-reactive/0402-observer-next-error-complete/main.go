package main

import "fmt"

// ObserverNextErrorComplete enforces the observer contract: a stream of next
// values terminated by exactly one complete or error. Once a terminal arrives,
// the "stopped" flag is set and all subsequent calls become no-ops.
type ObserverNextErrorComplete struct {
	stopped bool
	onNext  func(int)
}

func (o *ObserverNextErrorComplete) next(v int) {
	if o.stopped {
		return
	}
	o.onNext(v)
}

func (o *ObserverNextErrorComplete) complete() {
	if o.stopped {
		return
	}
	o.stopped = true
	fmt.Println("complete")
}

func (o *ObserverNextErrorComplete) error(err error) {
	if o.stopped {
		return
	}
	o.stopped = true
	fmt.Println("error:", err)
}

func main() {
	obs := &ObserverNextErrorComplete{
		onNext: func(v int) { fmt.Println(v) },
	}

	obs.next(1)
	obs.next(2)
	obs.complete()
	obs.next(3) // ignored: arrives after the terminal
}
