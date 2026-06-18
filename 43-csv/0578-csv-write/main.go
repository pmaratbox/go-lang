package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	w.WriteAll([][]string{
		{"name", "age"},
		{"Alice", "30"},
	})
	w.Flush()

	out := strings.ReplaceAll(buf.String(), "\r\n", "\n")
	out = strings.TrimRight(out, "\n")
	fmt.Println(out)
}
