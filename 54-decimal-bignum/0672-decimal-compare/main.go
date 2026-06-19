// Exact decimal comparison with math/big's Rat (exact rational) type.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Rat).SetString("0.1")
	b, _ := new(big.Rat).SetString("0.2")
	sum := new(big.Rat).Add(a, b)
	three, _ := new(big.Rat).SetString("0.3")
	fmt.Println(sum.Cmp(three) == 0) // true (exact, unlike float)
}
