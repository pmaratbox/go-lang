package main

import (
	"fmt"
	"strings"
)

func rwx(bits int) string {
	var b strings.Builder
	if bits&0b100 != 0 {
		b.WriteByte('r')
	} else {
		b.WriteByte('-')
	}
	if bits&0b010 != 0 {
		b.WriteByte('w')
	} else {
		b.WriteByte('-')
	}
	if bits&0b001 != 0 {
		b.WriteByte('x')
	} else {
		b.WriteByte('-')
	}
	return b.String()
}

func main() {
	fmt.Println(rwx(0b101))
}
