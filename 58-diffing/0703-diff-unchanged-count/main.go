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
	equal := 0
	for _, op := range sm.GetOpCodes() {
		switch op.Tag {
		case 'e': // equal: lines present unchanged in both
			equal += op.I2 - op.I1
		}
	}

	fmt.Println(equal)
}
