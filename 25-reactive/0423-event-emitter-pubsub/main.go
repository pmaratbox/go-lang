package main

import "fmt"

// handler is a topic subscriber identified by a comparable key so it can be removed.
type handler struct {
	id int
	fn func(payload string)
}

// EventEmitter is a multi-topic pub/sub registry.
type EventEmitter struct {
	topics map[string][]handler
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{topics: make(map[string][]handler)}
}

func (e *EventEmitter) on(topic string, h handler) {
	e.topics[topic] = append(e.topics[topic], h)
}

func (e *EventEmitter) emit(topic, payload string) {
	for _, h := range e.topics[topic] {
		h.fn(payload)
	}
}

func (e *EventEmitter) off(topic string, id int) {
	hs := e.topics[topic]
	kept := hs[:0]
	for _, h := range hs {
		if h.id != id {
			kept = append(kept, h)
		}
	}
	e.topics[topic] = kept
}

func main() {
	e := NewEventEmitter()

	h := handler{id: 1, fn: func(payload string) { fmt.Println("hi " + payload) }}
	g := handler{id: 2, fn: func(payload string) { fmt.Println("bye " + payload) }}

	e.on("greet", h)
	e.on("bye", g)

	e.emit("greet", "ada")
	e.emit("bye", "ada")

	e.off("greet", h.id)
	e.emit("greet", "x") // nothing: handler removed
}
