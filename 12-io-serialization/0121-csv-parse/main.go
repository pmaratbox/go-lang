package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	input := "alice,30\nbob,25\n"
	r := csv.NewReader(strings.NewReader(input))
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	pairs := make([]string, 0, len(records))
	for _, rec := range records {
		pairs = append(pairs, fmt.Sprintf("%s=%s", rec[0], rec[1]))
	}
	fmt.Println(strings.Join(pairs, " "))
}
