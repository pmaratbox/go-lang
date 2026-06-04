package main

import "fmt"

func main() {
	rows := [][2]string{{"a", "1"}, {"bb", "22"}}
	width := 0
	for _, r := range rows {
		if len(r[0]) > width {
			width = len(r[0])
		}
	}
	for _, r := range rows {
		fmt.Printf("%-*s | %s\n", width, r[0], r[1])
	}
}
