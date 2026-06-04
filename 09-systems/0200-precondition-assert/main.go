package main

import (
	"errors"
	"fmt"
)

func check(arg int) error {
	if arg > 0 {
		return nil
	}
	return errors.New("must be positive")
}

func main() {
	if err := check(5); err == nil {
		fmt.Println("ok")
	} else {
		fmt.Printf("error: %v\n", err)
	}

	if err := check(-1); err == nil {
		fmt.Println("ok")
	} else {
		fmt.Printf("error: %v\n", err)
	}
}
