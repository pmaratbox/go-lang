package main

import "fmt"

type person struct {
	id   int
	name string
}

func main() {
	people := []person{{1, "alice"}, {2, "bob"}}

	byID := make(map[int]string, len(people))
	for _, p := range people {
		byID[p.id] = p.name
	}

	fmt.Printf("id 2: %s\n", byID[2])
}
