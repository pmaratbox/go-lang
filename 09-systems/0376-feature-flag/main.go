package main

import "fmt"

func label(flag bool) string {
	if flag {
		return "enabled"
	}
	return "disabled"
}

func main() {
	fmt.Println(label(true), label(false))
}
