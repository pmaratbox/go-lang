package main

import "fmt"

// Signal is a writable reactive cell: it holds a value and a list of
// subscriber callbacks that are notified whenever the value changes.
type Signal struct {
	value       int
	subscribers []func()
}

func NewSignal(v int) *Signal {
	return &Signal{value: v}
}

// Get reads the current value.
func (s *Signal) Get() int {
	return s.value
}

// Set updates the value and notifies every subscriber.
func (s *Signal) Set(v int) {
	s.value = v
	for _, sub := range s.subscribers {
		sub()
	}
}

// Subscribe registers a recompute callback to run on every change.
func (s *Signal) Subscribe(fn func()) {
	s.subscribers = append(s.subscribers, fn)
}

// Computed caches a derived value, recomputing it whenever any signal it
// reads changes.
type Computed struct {
	cached  int
	compute func() int
}

// NewComputed wires the recompute callback as a subscriber of each dependency.
func NewComputed(compute func() int, deps ...*Signal) *Computed {
	c := &Computed{compute: compute}
	c.cached = compute()
	for _, dep := range deps {
		dep.Subscribe(func() { c.cached = c.compute() })
	}
	return c
}

// Get returns the cached derived value.
func (c *Computed) Get() int {
	return c.cached
}

func main() {
	a := NewSignal(2)
	b := NewSignal(3)

	sum := NewComputed(func() int { return a.Get() + b.Get() }, a, b)

	fmt.Println(sum.Get())

	a.Set(10)

	fmt.Println(sum.Get())
}
