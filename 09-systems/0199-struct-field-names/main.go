package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	t := reflect.TypeOf(Point{})
	var names []string
	for i := 0; i < t.NumField(); i++ {
		names = append(names, t.Field(i).Name)
	}
	fmt.Println(strings.Join(names, " "))
}
