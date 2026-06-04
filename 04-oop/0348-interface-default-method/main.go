package main

import "fmt"

type Greeter interface {
	greet() string
}

// Default embeds the default greet() returning "hi".
type Default struct{}

func (Default) greet() string { return "hi" }

// Override embeds Default but overrides greet() to "hey".
type Override struct {
	Default
}

func (Override) greet() string { return "hey" }

func main() {
	var d Greeter = Default{}
	var o Greeter = Override{}
	fmt.Println(d.greet(), o.greet())
}
