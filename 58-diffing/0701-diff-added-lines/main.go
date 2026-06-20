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
	var added []string
	for _, op := range sm.GetOpCodes() {
		switch op.Tag {
		case 'i':
			added = append(added, b[op.J1:op.J2]...)
		case 'r':
			added = append(added, b[op.J1:op.J2]...)
		}
	}

	fmt.Println(strings.Join(added, ","))
}
