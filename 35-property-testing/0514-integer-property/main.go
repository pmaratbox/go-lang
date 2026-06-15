// gopter — programmatic via Properties.Run with a discarding reporter (no testing.T needed).
// Run returns true if all properties pass. Use io.Discard so nothing leaks to stdout.
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
	props.Property("addition-commutative", prop.ForAll(func(a, b int) bool {
		return a+b == b+a
	}, gen.Int(), gen.Int()))
	if props.Run(gopter.NewFormatedReporter(false, 80, io.Discard)) {
		fmt.Println("passed")
	}
}
