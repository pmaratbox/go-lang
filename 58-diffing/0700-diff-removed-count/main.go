// Go — pmezard/go-difflib (line-oriented SequenceMatcher with opcodes).
package main

import (
	"fmt"

	"github.com/pmezard/go-difflib/difflib"
)

func main() {
	a := []string{"apple", "banana", "cherry"}
	b := []string{"apple", "blueberry", "cherry", "date"}

	sm := difflib.NewMatcher(a, b)
	var removed []string
	for _, op := range sm.GetOpCodes() {
		switch op.Tag {
		case 'd': // delete: source-only lines
			removed = append(removed, a[op.I1:op.I2]...)
		case 'r': // replace: maps to removed(source) + added(target)
			removed = append(removed, a[op.I1:op.I2]...)
		}
	}

	fmt.Println(len(removed))
}
