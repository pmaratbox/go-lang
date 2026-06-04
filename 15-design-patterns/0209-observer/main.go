package main

import "fmt"

type observer interface {
	update(value int)
}

type namedObserver struct{ id string }

func (o namedObserver) update(value int) {
	fmt.Printf("%s: %d\n", o.id, value)
}

type subject struct{ observers []observer }

func (s *subject) register(o observer) {
	s.observers = append(s.observers, o)
}

func (s *subject) notify(value int) {
	for _, o := range s.observers {
		o.update(value)
	}
}

func main() {
	s := &subject{}
	s.register(namedObserver{id: "obs1"})
	s.register(namedObserver{id: "obs2"})
	s.notify(5)
}
