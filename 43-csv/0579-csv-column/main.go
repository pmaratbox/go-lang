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
	// Find the "age" column index from the header row.
	header := rows[0]
	col := -1
	for i, name := range header {
		if name == "age" {
			col = i
		}
	}
	// Collect the age value from each data row.
	ages := []string{}
	for _, r := range rows[1:] {
		ages = append(ages, r[col])
	}
	fmt.Println(strings.Join(ages, ","))
}
