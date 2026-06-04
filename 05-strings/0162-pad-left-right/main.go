package main

import "fmt"

func main() {
	left := fmt.Sprintf("%3s", "5")
	right := fmt.Sprintf("%-3s", "5")
	fmt.Println(left + "|" + right)
}
