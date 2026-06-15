// gopter — programmatic via Properties.Run with a discarding reporter (no testing.T needed).
// A FALSE property: "every non-negative integer is < 100". gopter generates
// inputs, finds a counterexample, and Run returns false. The falsifying-example
// and shrink report is sent to io.Discard so ONLY `found` reaches stdout.
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
	props.Property("every-non-negative-int-lt-100", prop.ForAll(func(n int) bool {
		return n < 100
	}, gen.IntRange(0, 1000000)))
	// Discard the reporter so the falsifying-example/shrink report never hits stdout.
	if props.Run(gopter.NewFormatedReporter(false, 80, io.Discard)) {
		fmt.Println("no-counterexample")
	} else {
		fmt.Println("found")
	}
}
