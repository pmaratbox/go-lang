// gopter — programmatic via Properties.Run with a discarding reporter (no testing.T needed).
// A CUSTOM generator built with the Map combinator: gen.Int().Map(n -> n*2) yields even ints.
package main

import (
	"fmt"
	"io"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func main() {
	props := gopter.NewProperties(gopter.DefaultTestParametersWithSeed(42))

	// Custom generator: map each generated int n to 2*n, so every value is even.
	evenGen := gen.Int().Map(func(n int) int { return n * 2 })

	props.Property("always-even", prop.ForAll(func(n int) bool {
		return n%2 == 0
	}, evenGen))

	if props.Run(gopter.NewFormatedReporter(false, 80, io.Discard)) {
		fmt.Println("passed")
	}
}
