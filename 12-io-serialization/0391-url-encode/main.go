package main

import (
	"fmt"
	"strings"
)

func isUnreserved(c byte) bool {
	switch {
	case c >= 'A' && c <= 'Z',
		c >= 'a' && c <= 'z',
		c >= '0' && c <= '9':
		return true
	case c == '-' || c == '_' || c == '.' || c == '~':
		return true
	}
	return false
}

func main() {
	input := "a b&c"

	var sb strings.Builder
	for i := 0; i < len(input); i++ {
		c := input[i]
		if isUnreserved(c) {
			sb.WriteByte(c)
		} else {
			fmt.Fprintf(&sb, "%%%02X", c)
		}
	}
	fmt.Println(sb.String())
}
