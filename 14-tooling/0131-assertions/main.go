package main

import "fmt"

func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func main() {
	assert(1+1 == 2, "1+1 should be 2")
	assert(2*3 == 6, "2*3 should be 6")
	assert("go"+"lang" == "golang", "concat should be golang")
	fmt.Println("all passed")
}
