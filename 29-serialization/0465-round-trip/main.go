package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func main() {
	b, err := json.Marshal(Person{Age: 30, Name: "alice"})
	if err != nil {
		panic(err)
	}
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		panic(err)
	}
	fmt.Println(p.Name)
}
