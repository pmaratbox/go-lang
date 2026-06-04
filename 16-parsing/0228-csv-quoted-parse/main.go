package main

import (
	"fmt"
	"strings"
)

func main() {
	row := `a,"b,c",d`
	var fields []string
	var cur strings.Builder
	inQuotes := false
	for _, ch := range row {
		switch {
		case ch == '"':
			inQuotes = !inQuotes
		case ch == ',' && !inQuotes:
			fields = append(fields, cur.String())
			cur.Reset()
		default:
			cur.WriteRune(ch)
		}
	}
	fields = append(fields, cur.String())
	fmt.Println(strings.Join(fields, "|"))
}
