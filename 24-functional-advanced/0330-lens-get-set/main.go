package main

import "fmt"

type inner struct{ b int }
type outer struct{ a inner }

// Lens pairs a getter with an immutable setter.
type Lens struct {
	get func(outer) int
	set func(outer, int) outer
}

func main() {
	bLens := Lens{
		get: func(o outer) int { return o.a.b },
		set: func(o outer, v int) outer {
			o.a.b = v // o is a value copy, so the original is untouched
			return o
		},
	}

	original := outer{a: inner{b: 1}}
	got := bLens.get(original)
	updated := bLens.set(original, 2)
	fmt.Println(got, bLens.get(updated))
}
