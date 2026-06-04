package main

import "fmt"

// thunk represents a suspended computation: either a final value or more work.
type thunk struct {
	done bool
	val  int
	next func() thunk
}

func sumTo(n, acc int) thunk {
	if n == 0 {
		return thunk{done: true, val: acc}
	}
	return thunk{next: func() thunk { return sumTo(n-1, acc+n) }}
}

// trampoline drives thunks in a flat loop, avoiding deep recursion.
func trampoline(t thunk) int {
	for !t.done {
		t = t.next()
	}
	return t.val
}

func main() {
	fmt.Println(trampoline(sumTo(100, 0)))
}
