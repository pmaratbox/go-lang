package main

import (
	"fmt"

	"github.com/pmezard/go-difflib/difflib"
)

func main() {
	a := []string{"apple", "banana", "cherry"}
	b := []string{"apple", "blueberry", "cherry", "date"}

	sm := difflib.NewMatcher(a, b)
	var added, removed []string
	unchanged := 0
	for _, op := range sm.GetOpCodes() {
		switch op.Tag {
		case 'i': // insert: lines present only in b
			added = append(added, b[op.J1:op.J2]...)
		case 'd': // delete: lines present only in a
			removed = append(removed, a[op.I1:op.I2]...)
		case 'r': // replace: maps to removed(source)+added(target)
			removed = append(removed, a[op.I1:op.I2]...)
			added = append(added, b[op.J1:op.J2]...)
		case 'e': // equal: unchanged lines
			unchanged += op.I2 - op.I1
		}
	}

	fmt.Printf("%d %d %d\n", len(added), len(removed), unchanged)
}
