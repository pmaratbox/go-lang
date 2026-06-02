package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "world"

	fmt.Printf("Hello, %s!\n", name)
	fmt.Println(strings.ToUpper(name))
	fmt.Printf("length: %d\n", len(name))
}
