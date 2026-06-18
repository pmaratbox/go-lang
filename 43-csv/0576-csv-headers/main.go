package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	data := "name,age,city\nAlice,30,Paris\nBob,25,London\nCarol,35,Berlin\n"
	rows, err := csv.NewReader(strings.NewReader(data)).ReadAll()
	if err != nil {
		panic(err)
	}
	// rows[0] is the header row; join its fields with a pipe.
	fmt.Println(strings.Join(rows[0], "|"))
}
