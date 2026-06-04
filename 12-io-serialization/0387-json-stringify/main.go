package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	obj := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{Name: "Ada", Age: 36}

	b, _ := json.Marshal(obj)
	fmt.Println(string(b))
}
