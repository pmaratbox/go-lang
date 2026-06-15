// gopter — programmatic property check via Properties.Run with a discarding reporter (no testing.T).
// Run returns true only if all properties pass. io.Discard keeps any report off stdout.
package main

import (
	"fmt"
	"io"
	"sort"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func sorted(xs []int) []int {
	r := make([]int, len(xs))
	copy(r, xs)
	sort.Ints(r)
	return r
}

func main() {
	props := gopter.NewProperties(gopter.DefaultTestParametersWithSeed(42))
	props.Property("sort-idempotent", prop.ForAll(func(xs []int) bool {
		once := sorted(xs)
		twice := sorted(once)
		for i := range once {
			if once[i] != twice[i] {
				return false
			}
		}
		return true
	}, gen.SliceOf(gen.Int())))

	if props.Run(gopter.NewFormatedReporter(false, 80, io.Discard)) {
		fmt.Println("passed")
	}
}
