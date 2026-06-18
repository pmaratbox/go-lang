package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	// Build a map with string and integer values.
	m := map[string]any{
		"name": "Alice",
		"age":  30,
		"city": "Paris",
	}

	// yaml.Marshal serializes a map[string]any with SORTED keys in
	// block style (no flow braces, no quotes on simple scalars).
	out, err := yaml.Marshal(m)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(out)
}
