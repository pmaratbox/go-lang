package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Compute 30! exactly with arbitrary-precision big.Int.
	f := big.NewInt(1)
	for i := int64(1); i <= 30; i++ {
		f.Mul(f, big.NewInt(i))
	}
	fmt.Println(f)
}
