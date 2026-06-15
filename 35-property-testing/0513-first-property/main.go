// gopter — programmatic property check via Properties.Run with a discarding reporter (no testing.T).
// Run returns true only if all properties pass. io.Discard keeps any report off stdout.
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
	props.Property("reverse-twice", prop.ForAll(func(xs []int) bool {
		r := make([]int, len(xs))
		copy(r, xs)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		for i := range xs {
			if r[i] != xs[i] {
				return false
			}
		}
		return true
	}, gen.SliceOf(gen.Int())))

	if props.Run(gopter.NewFormatedReporter(false, 80, io.Discard)) {
		fmt.Println("passed")
	}
}
