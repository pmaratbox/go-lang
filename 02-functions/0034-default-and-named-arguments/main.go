package main

import "fmt"

func greet(name, greeting string) string {
	return greeting + ", " + name
}

func greetDefault(name string) string {
	return greet(name, "Hello")
}

func main() {
	fmt.Println(greetDefault("Ada"))
	fmt.Println(greet("Ada", "Hi"))
}
