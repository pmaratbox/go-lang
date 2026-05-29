package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", result)
	}

	if _, err := divide(10, 0); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
