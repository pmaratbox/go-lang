package main

import "fmt"

func isPalindrome(s string) bool {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

func main() {
	for _, word := range []string{"level", "hello"} {
		ans := "no"
		if isPalindrome(word) {
			ans = "yes"
		}
		fmt.Printf("%s: %s\n", word, ans)
	}
}
