package main

import (
	"fmt"
	"math/big"
)

func main() {
	// big.Rat is an exact rational, so 1.1 * 1.1 is computed with no
	// floating-point rounding. FloatString(2) renders the two decimal places.
	a, _ := new(big.Rat).SetString("1.1")
	b, _ := new(big.Rat).SetString("1.1")
	product := new(big.Rat).Mul(a, b)
	fmt.Println(product.FloatString(2))
}
