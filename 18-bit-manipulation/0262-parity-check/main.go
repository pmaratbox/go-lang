package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(bits.OnesCount(7)&1, bits.OnesCount(5)&1)
}
