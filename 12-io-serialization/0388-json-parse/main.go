package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := `{"x":1,"y":2}`

	// Preserve key order explicitly for stable output.
	var m map[string]int
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		panic(err)
	}

	keys := []string{"x", "y"}
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%d", k, m[k]))
	}
	fmt.Println(parts[0] + " " + parts[1])
}
