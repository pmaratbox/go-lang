package main

import (
	"fmt"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
)

func main() {
	a := []string{"apple", "banana", "cherry"}
	b := []string{"apple", "blueberry", "cherry", "date"}

	sm := difflib.NewMatcher(a, b)
	var removed []string
	for _, op := range sm.GetOpCodes() {
		switch op.Tag {
		case 'd':
			removed = append(removed, a[op.I1:op.I2]...)
		case 'r':
			removed = append(removed, a[op.I1:op.I2]...)
		}
	}

	fmt.Println(strings.Join(removed, ","))
}
