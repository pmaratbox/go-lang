// Exact decimal subtraction with math/big's Rat (exact rational) type.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Rat).SetString("1.0")
	b, _ := new(big.Rat).SetString("0.1")
	diff := new(big.Rat).Sub(a, b)
	fmt.Println(diff.FloatString(1)) // 0.9
}
