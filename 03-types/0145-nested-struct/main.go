package main

import "fmt"

type Address struct {
	City string
}

type Person struct {
	Name    string
	Address Address
}

func main() {
	p := Person{
		Name:    "Ada",
		Address: Address{City: "London"},
	}

	fmt.Println(p.Address.City)
}
