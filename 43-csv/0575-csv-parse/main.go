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
	names := []string{}
	for _, r := range rows[1:] { // skip header row
		names = append(names, r[0]) // first column = name
	}
	fmt.Println(strings.Join(names, ","))
}
