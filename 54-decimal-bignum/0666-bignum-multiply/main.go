package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(123456789)
	b := big.NewInt(987654321)
	product := new(big.Int).Mul(a, b)
	fmt.Println(product)
}
