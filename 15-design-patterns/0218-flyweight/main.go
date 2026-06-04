package main

import "fmt"

type glyph struct{ char string }

type factory struct{ cache map[string]*glyph }

func newFactory() *factory { return &factory{cache: map[string]*glyph{}} }

func (f *factory) get(char string) *glyph {
	if g, ok := f.cache[char]; ok {
		return g
	}
	g := &glyph{char: char}
	f.cache[char] = g
	return g
}

func main() {
	f := newFactory()
	for _, c := range []string{"a", "b", "a"} {
		f.get(c)
	}
	fmt.Println(len(f.cache))
}
