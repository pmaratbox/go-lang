package main

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "name,age,city\nAlice,30,Paris\nBob,25,London\nCarol,35,Berlin\n"
	rows, _ := csv.NewReader(strings.NewReader(data)).ReadAll()

	names := []string{}
	for _, r := range rows[1:] {
		age, _ := strconv.Atoi(r[1])
		if age > 28 {
			names = append(names, r[0])
		}
	}
	fmt.Println(strings.Join(names, ","))
}
