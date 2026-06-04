package main

import (
	"errors"
	"fmt"
)

func main() {
	inner := errors.New("inner")
	outer := fmt.Errorf("outer: %w", inner)
	fmt.Println(outer)
}
