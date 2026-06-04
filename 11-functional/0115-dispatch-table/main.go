package main

import "fmt"

func main() {
	ops := map[string]func(int, int) int{
		"add": func(a, b int) int { return a + b },
		"mul": func(a, b int) int { return a * b },
	}
	fmt.Println(ops["add"](3, 4), ops["mul"](3, 4))
}
