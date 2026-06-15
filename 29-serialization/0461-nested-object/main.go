package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City string `json:"city"`
	Zip  int    `json:"zip"`
}

type Person struct {
	Address Address `json:"address"`
	Name    string  `json:"name"`
}

func main() {
	p := Person{Address: Address{City: "oslo", Zip: 1000}, Name: "alice"}
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
