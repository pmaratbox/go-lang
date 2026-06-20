package main

import (
	"fmt"

	"github.com/pmezard/go-difflib/difflib"
)

func main() {
	a := []string{"apple", "banana", "cherry"}
	b := []string{"apple", "blueberry", "cherry", "date"}

	sm := difflib.NewMatcher(a, b)
	var added []string
	for _, op := range sm.GetOpCodes() {
		switch op.Tag {
		case 'i': // insert: lines present only in b
			added = append(added, b[op.J1:op.J2]...)
		case 'r': // replace: maps to removed(source)+added(target)
			added = append(added, b[op.J1:op.J2]...)
		}
	}

	fmt.Println(len(added))
}
