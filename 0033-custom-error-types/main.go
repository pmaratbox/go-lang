package main

import "fmt"

type TooLargeError struct{}

func (e *TooLargeError) Error() string {
	return "value too large"
}

func check(n int) error {
	if n > 100 {
		return &TooLargeError{}
	}
	return nil
}

func main() {
	if err := check(200); err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
