package main

import (
	"encoding/json"
	"fmt"
)

// The struct field is FullName, but the json struct tag renames it to the
// JSON key "full_name".
type Person struct {
	FullName string `json:"full_name"`
}

func main() {
	b, err := json.Marshal(Person{FullName: "alice"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
