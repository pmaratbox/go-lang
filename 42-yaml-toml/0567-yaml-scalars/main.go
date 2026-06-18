package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func main() {
	var d map[string]any
	yaml.Unmarshal([]byte("name: Alice\nrole: admin\nage: 30\n"), &d)
	fmt.Printf("%v %v %v\n", d["name"], d["role"], d["age"])
}
