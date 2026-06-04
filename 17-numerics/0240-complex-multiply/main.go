package main

import "fmt"

func main() {
	z := complex(1, 2) * complex(3, 4)
	fmt.Printf("%g %g\n", real(z), imag(z))
}
