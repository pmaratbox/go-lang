// gopter — programmatic via Properties.Run with a discarding reporter (no testing.T needed).
// A property over TWO generated integer arguments a, b.
package main

import (
	"fmt"
	"io"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	props := gopter.NewProperties(gopter.DefaultTestParametersWithSeed(42))
	props.Property("max-dominates-both", prop.ForAll(func(a, b int) bool {
		m := maxInt(a, b)
		return m >= a && m >= b
	}, gen.Int(), gen.Int()))
	if props.Run(gopter.NewFormatedReporter(false, 80, io.Discard)) {
		fmt.Println("passed")
	}
}
