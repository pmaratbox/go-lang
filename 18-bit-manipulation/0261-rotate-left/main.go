package main

import "fmt"

func rol8(x byte, n int) byte {
	return byte((uint(x)<<n | uint(x)>>(8-n)) & 0xff)
}

func main() {
	fmt.Println(rol8(1, 1), rol8(128, 1))
}
