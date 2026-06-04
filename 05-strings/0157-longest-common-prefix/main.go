package main

import "fmt"

func main() {
	words := []string{"flower", "flow", "flight"}
	prefix := words[0]
	for _, w := range words[1:] {
		i := 0
		for i < len(prefix) && i < len(w) && prefix[i] == w[i] {
			i++
		}
		prefix = prefix[:i]
	}
	fmt.Println(prefix)
}
