package main

import (
	"errors"
	"fmt"
)

// baseError is the supertype; specificError wraps it as a subtype.
var baseError = errors.New("base")

type specificError struct{}

func (specificError) Error() string { return "specific" }
func (specificError) Unwrap() error { return baseError }

func main() {
	var err error = specificError{}
	if errors.Is(err, baseError) {
		fmt.Println("caught base")
	}
}
