package main

import (
	"fmt"
	"math/big"
)

func main() {
	// big.Rat is an exact rational, so 0.1 and 0.2 are stored with no
	// floating-point rounding error.
	a, _ := new(big.Rat).SetString("0.1")
	b, _ := new(big.Rat).SetString("0.2")

	sum := new(big.Rat).Add(a, b)

	// FloatString(1) renders the exact value with one decimal place.
	fmt.Println(sum.FloatString(1))
}
