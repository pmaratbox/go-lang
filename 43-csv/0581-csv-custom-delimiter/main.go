package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	data := "a;b;c\n1;2;3\n"
	rd := csv.NewReader(strings.NewReader(data))
	rd.Comma = ';'
	rows, _ := rd.ReadAll()
	fmt.Println(strings.Join(rows[1], ","))
}
