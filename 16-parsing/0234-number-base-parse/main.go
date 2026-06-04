package main

import (
	"fmt"
	"strconv"
)

func main() {
	hex, _ := strconv.ParseInt("ff", 16, 64)
	bin, _ := strconv.ParseInt("101", 2, 64)
	fmt.Println(hex, bin)
}
