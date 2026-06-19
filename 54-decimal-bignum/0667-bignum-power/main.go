package main

import (
	"fmt"
	"math/big"
)

func main() {
	// big.Int is Go's arbitrary-precision integer. Exp computes
	// base**exponent exactly: here 2^100 with no float rounding.
	result := new(big.Int).Exp(big.NewInt(2), big.NewInt(100), nil)
	fmt.Println(result)
}
