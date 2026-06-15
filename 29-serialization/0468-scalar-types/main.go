package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Active bool   `json:"active"`
	Count  int    `json:"count"`
	Label  string `json:"label"`
}

func main() {
	b, _ := json.Marshal(Item{Active: true, Count: 5, Label: "x"})
	fmt.Println(string(b))
}
