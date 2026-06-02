package main

import "fmt"

type Animal struct{}

func (a Animal) Speak() string {
	return "some sound"
}

type Dog struct {
	Animal
}

func (d Dog) Speak() string {
	return "Woof"
}

func main() {
	fmt.Println("animal:", Animal{}.Speak())
	fmt.Println("dog:", Dog{}.Speak())
}
