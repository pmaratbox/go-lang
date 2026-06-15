// gopter — programmatic property check via Properties.Run with a discarding reporter (no testing.T).
// A precondition (SuchThat) filters the generator to positive integers before the predicate runs.
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

	// Precondition: only test positive integers. gen.Int().SuchThat keeps the
	// values that satisfy the predicate, so gopter generates positive n only.
	positive := gen.Int().SuchThat(func(n int) bool { return n > 0 })

	props.Property("n+1 > n for positive n", prop.ForAll(func(n int) bool {
		return n+1 > n
	}, positive))

	if props.Run(gopter.NewFormatedReporter(false, 80, io.Discard)) {
		fmt.Println("passed")
	}
}
