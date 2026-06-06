package main

import "fmt"

// Observer receives pushed values.
type Observer func(v int)

// BehaviorSubject holds a current value and replays it to each new subscriber.
type BehaviorSubject struct {
	current   int
	observers []Observer
}

// NewBehaviorSubject seeds the subject with an initial current value.
func NewBehaviorSubject(seed int) *BehaviorSubject {
	return &BehaviorSubject{current: seed}
}

// Subscribe registers an observer and immediately replays the current value.
func (s *BehaviorSubject) Subscribe(obs Observer) {
	s.observers = append(s.observers, obs)
	obs(s.current)
}

// Next updates the current value and pushes it to all observers.
func (s *BehaviorSubject) Next(v int) {
	s.current = v
	for _, obs := range s.observers {
		obs(v)
	}
}

func main() {
	subject := NewBehaviorSubject(0)

	subject.Subscribe(func(v int) { fmt.Printf("A: %d\n", v) })
	subject.Next(1)
	subject.Subscribe(func(v int) { fmt.Printf("B: %d\n", v) })
	subject.Next(2)
}
