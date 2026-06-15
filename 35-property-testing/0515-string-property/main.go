// gopter — programmatic property check via Properties.Run with a discarding reporter (no testing.T).
// Run returns true if all generated cases satisfy the property; io.Discard keeps stdout clean.
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
	props.Property("len(s+s) == 2*len(s)", prop.ForAll(func(s string) bool {
		return len(s+s) == 2*len(s)
	}, gen.AlphaString()))

	if props.Run(gopter.NewFormatedReporter(false, 80, io.Discard)) {
		fmt.Println("passed")
	}
}
