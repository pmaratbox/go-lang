package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	name := "data.txt"
	defer os.Remove(name)

	if err := os.WriteFile(name, []byte("a\n"), 0o644); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	if _, err := f.WriteString("b\n"); err != nil {
		panic(err)
	}
	f.Close()

	in, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	var lines []string
	sc := bufio.NewScanner(in)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	fmt.Println(strings.Join(lines, " "))
}
