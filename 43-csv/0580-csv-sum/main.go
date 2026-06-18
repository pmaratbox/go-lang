package main

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "name,age,city\nAlice,30,Paris\nBob,25,London\nCarol,35,Berlin\n"
	rows, err := csv.NewReader(strings.NewReader(data)).ReadAll()
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, r := range rows[1:] { // skip header row
		n, err := strconv.Atoi(r[1]) // age column
		if err != nil {
			panic(err)
		}
		sum += n
	}
	fmt.Println(sum)
}
