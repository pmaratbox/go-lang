package main

import (
	"encoding/json"
	"fmt"
)

// Fields are declared alphabetically so the encoded keys come out
// in alphabetical order (encoding/json preserves declaration order).
type Person struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func main() {
	p := Person{Age: 30, Name: "alice"}
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
