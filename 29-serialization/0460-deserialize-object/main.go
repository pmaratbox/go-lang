package main

import (
	"encoding/json"
	"fmt"
)

// Fields declared alphabetically so the canonical compact JSON has alphabetical keys.
type Person struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func main() {
	data := []byte(`{"age":30,"name":"alice"}`)

	var p Person
	if err := json.Unmarshal(data, &p); err != nil {
		panic(err)
	}

	fmt.Println(p.Name, p.Age)
}
