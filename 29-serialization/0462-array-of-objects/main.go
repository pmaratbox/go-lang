package main

import (
	"encoding/json"
	"fmt"
)

// Person fields are declared alphabetically so json.Marshal emits keys
// (age, name) in alphabetical order.
type Person struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func main() {
	people := []Person{
		{Age: 30, Name: "alice"},
		{Age: 25, Name: "bob"},
	}
	b, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
