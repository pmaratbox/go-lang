package main

import (
	"fmt"
	"sort"
	"strings"
)

type person struct {
	name string
	age  int
}

func main() {
	people := []person{{"alice", 30}, {"bob", 25}}
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})

	names := make([]string, len(people))
	for i, p := range people {
		names[i] = p.name
	}
	fmt.Println(strings.Join(names, " "))
}
