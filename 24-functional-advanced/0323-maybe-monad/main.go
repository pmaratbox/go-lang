package main

import "fmt"

// Maybe is an optional int value.
type Maybe struct {
	value   int
	present bool
}

func Some(v int) Maybe { return Maybe{value: v, present: true} }
func None() Maybe      { return Maybe{} }

// bind applies f only when a value is present.
func (m Maybe) bind(f func(int) Maybe) Maybe {
	if !m.present {
		return m
	}
	return f(m.value)
}

func main() {
	add3 := func(x int) Maybe { return Some(x + 3) }
	mul2 := func(x int) Maybe { return Some(x * 2) }

	present := Some(2).bind(add3).bind(mul2)
	absent := None().bind(add3).bind(mul2)

	left := fmt.Sprintf("%d", present.value)
	right := "none"
	if absent.present {
		right = fmt.Sprintf("%d", absent.value)
	}
	fmt.Println(left, right)
}
