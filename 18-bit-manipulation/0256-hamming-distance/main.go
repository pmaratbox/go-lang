package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(bits.OnesCount(uint(1 ^ 4)))
}
