package main

import (
	"encoding/json"
	"fmt"
)

// Fields declared alphabetically so the marshaled JSON keys stay in
// alphabetical order. Age has no struct tag option for a default, but an
// absent JSON field simply leaves Go's zero value (0) in place.
type Person struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func main() {
	input := []byte(`{"name":"alice"}`)

	var p Person
	if err := json.Unmarshal(input, &p); err != nil {
		panic(err)
	}

	fmt.Println(p.Name, p.Age)
}
