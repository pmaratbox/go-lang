package main

import (
	"fmt"

	"github.com/pmezard/go-difflib/difflib"
)

func main() {
	a := []string{"apple", "banana", "cherry"}
	b := []string{"apple", "banana", "cherry"}

	sm := difflib.NewMatcher(a, b)
	added, removed := 0, 0
	for _, op := range sm.GetOpCodes() {
		switch op.Tag {
		case 'i':
			added += op.J2 - op.J1
		case 'd':
			removed += op.I2 - op.I1
		case 'r':
			added += op.J2 - op.J1
			removed += op.I2 - op.I1
		}
	}

	fmt.Printf("%d %d\n", added, removed)
}
