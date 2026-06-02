package main

import "fmt"

func main() {
	text := "abc"
	shift := 1
	result := []rune{}
	for _, ch := range text {
		shifted := (ch-'a'+rune(shift))%26 + 'a'
		result = append(result, shifted)
	}
	fmt.Println(string(result))
}
