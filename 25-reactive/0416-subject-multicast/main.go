package main

import "fmt"

// Observer receives pushed values via its next callback.
type Observer struct {
	next func(int)
}

// Subject multicasts each emission to all current observers.
type Subject struct {
	observers []*Observer
}

func (s *Subject) subscribe(o *Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) emit(v int) {
	for _, o := range s.observers {
		o.next(v)
	}
}

func main() {
	subject := &Subject{}

	subject.subscribe(&Observer{next: func(v int) {
		fmt.Println("obs1:", v)
	}})
	subject.subscribe(&Observer{next: func(v int) {
		fmt.Println("obs2:", v)
	}})

	subject.emit(1)
	subject.emit(2)
}
