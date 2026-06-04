package main

import "fmt"

type Result struct {
	value int
	err   string
}

func ok(v int) Result    { return Result{value: v} }
func errf(m string) Result { return Result{err: m} }

func safeDiv(a, b int) Result {
	if b == 0 {
		return errf("divide by zero")
	}
	return ok(a / b)
}

func (r Result) String() string {
	if r.err != "" {
		return "err: " + r.err
	}
	return fmt.Sprintf("ok: %d", r.value)
}

func main() {
	fmt.Println(safeDiv(10, 2))
	fmt.Println(safeDiv(1, 0))
}
