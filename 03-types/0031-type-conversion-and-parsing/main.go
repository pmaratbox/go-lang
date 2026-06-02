package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi("42")
	f, _ := strconv.ParseFloat("3.5", 64)
	s := strconv.Itoa(n)
	fmt.Printf("int: %d\n", n)
	fmt.Printf("float: %g\n", f)
	fmt.Printf("str: %s\n", s)
}
