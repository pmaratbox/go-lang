package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("sqrt:", int(math.Sqrt(16)))
	fmt.Println("pow:", int(math.Pow(2, 10)))
	fmt.Println("abs:", int(math.Abs(-5)))
	fmt.Println("max:", max(3, 9))
}
