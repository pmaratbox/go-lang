package main

import "fmt"

func main() {
	messages := map[int]string{
		0: "ok",
		1: "denied",
		2: "not found",
	}
	fmt.Println(messages[2])
}
