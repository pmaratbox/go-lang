package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	data := "name,note\nAlice,\"hello, world\"\n"
	rows, err := csv.NewReader(strings.NewReader(data)).ReadAll()
	if err != nil {
		panic(err)
	}
	// rows[0] is the header (name,note); rows[1] is the data row.
	// The quoted field "hello, world" keeps its embedded comma intact.
	fmt.Println(rows[1][1]) // note column of the data row
}
