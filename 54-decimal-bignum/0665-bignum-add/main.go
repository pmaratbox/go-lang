package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("12345678901234567890", 10)
	b, _ := new(big.Int).SetString("98765432109876543210", 10)
	fmt.Println(new(big.Int).Add(a, b))
}
