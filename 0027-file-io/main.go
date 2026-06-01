package main

import (
	"fmt"
	"os"
)

func main() {
	path := "greeting.txt"
	if err := os.WriteFile(path, []byte("hello, file"), 0644); err != nil {
		panic(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	os.Remove(path)
	fmt.Printf("read: %s\n", data)
}
