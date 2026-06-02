package main

import "fmt"

func main() {
	text := "aaabbc"
	runes := []rune(text)
	i := 0
	for i < len(runes) {
		ch := runes[i]
		count := 0
		for i < len(runes) && runes[i] == ch {
			count++
			i++
		}
		fmt.Printf("%c%d", ch, count)
	}
	fmt.Println()
}
