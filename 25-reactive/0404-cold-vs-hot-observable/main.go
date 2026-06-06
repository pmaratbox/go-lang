package main

import (
	"fmt"
	"strings"
)

// observer receives pushed values.
type observer struct {
	next func(int)
}

// coldObservable re-runs its producer for each subscriber.
type coldObservable struct {
	producer func(observer)
}

func (c coldObservable) subscribe(o observer) {
	c.producer(o)
}

// hotObservable shares a single producer execution across subscribers.
// Late subscribers only see values emitted after they subscribe.
type hotObservable struct {
	observers *[]observer
}

func newHot() hotObservable {
	return hotObservable{observers: &[]observer{}}
}

func (h hotObservable) subscribe(o observer) {
	*h.observers = append(*h.observers, o)
}

func (h hotObservable) emit(v int) {
	for _, o := range *h.observers {
		o.next(v)
	}
}

func join(vs []int) string {
	parts := make([]string, len(vs))
	for i, v := range vs {
		parts[i] = fmt.Sprint(v)
	}
	return strings.Join(parts, " ")
}

func main() {
	// COLD: each subscription independently runs the producer.
	cold := coldObservable{producer: func(o observer) {
		for _, v := range []int{1, 2, 3} {
			o.next(v)
		}
	}}

	var coldA, coldB []int
	cold.subscribe(observer{next: func(v int) { coldA = append(coldA, v) }})
	cold.subscribe(observer{next: func(v int) { coldB = append(coldB, v) }})

	fmt.Println("cold A:", join(coldA))
	fmt.Println("cold B:", join(coldB))

	// HOT: one shared producer; A subscribes, emit 1, B subscribes, emit 2 3.
	hot := newHot()
	var hotA, hotB []int
	hot.subscribe(observer{next: func(v int) { hotA = append(hotA, v) }})
	hot.emit(1)
	hot.subscribe(observer{next: func(v int) { hotB = append(hotB, v) }})
	hot.emit(2)
	hot.emit(3)

	fmt.Println("hot A:", join(hotA))
	fmt.Println("hot B:", join(hotB))
}
