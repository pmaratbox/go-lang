package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Ada", Age: 36}

	fmt.Printf("name: %s\n", p.Name)
	fmt.Printf("age: %d\n", p.Age)
}
