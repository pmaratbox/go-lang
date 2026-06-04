package main

import (
	"errors"
	"fmt"
)

func main() {
	defer fmt.Println("cleanup")

	err := errors.New("boom")
	if err != nil {
		fmt.Println("caught")
	}
}
