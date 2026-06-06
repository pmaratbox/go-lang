package main

import "fmt"

// Observer receives pushed values.
type Observer func(int)

// ReplaySubject buffers the last bufSize emitted values and replays them to
// late subscribers, who then continue to receive new emissions.
type ReplaySubject struct {
	bufSize     int
	buffer      []int
	subscribers []Observer
}

func NewReplaySubject(bufSize int) *ReplaySubject {
	return &ReplaySubject{bufSize: bufSize}
}

// Next emits a value: it is buffered (trimmed to bufSize) and pushed to all
// current subscribers.
func (s *ReplaySubject) Next(v int) {
	s.buffer = append(s.buffer, v)
	if len(s.buffer) > s.bufSize {
		s.buffer = s.buffer[len(s.buffer)-s.bufSize:]
	}
	for _, sub := range s.subscribers {
		sub(v)
	}
}

// Subscribe replays the buffered values to the new observer, then registers it
// for future emissions.
func (s *ReplaySubject) Subscribe(obs Observer) {
	for _, v := range s.buffer {
		obs(v)
	}
	s.subscribers = append(s.subscribers, obs)
}

func main() {
	subject := NewReplaySubject(2)

	subject.Next(1)
	subject.Next(2)
	subject.Next(3) // buffer now [2, 3]

	// Late subscriber: immediately receives buffered 2 then 3.
	subject.Subscribe(func(v int) {
		fmt.Println(v)
	})

	subject.Next(4) // subscriber receives 4
}
