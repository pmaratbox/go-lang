package main

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	doc := "fruits:\n  - apple\n  - banana\n  - cherry\n"

	var d map[string]any
	if err := yaml.Unmarshal([]byte(doc), &d); err != nil {
		panic(err)
	}

	var fruits []string
	for _, v := range d["fruits"].([]any) {
		fruits = append(fruits, v.(string))
	}

	fmt.Println(strings.Join(fruits, ","))
}
