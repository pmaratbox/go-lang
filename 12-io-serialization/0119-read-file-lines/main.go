package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	name := "lines.txt"
	if err := os.WriteFile(name, []byte("one\ntwo\nthree\n"), 0o644); err != nil {
		panic(err)
	}
	defer os.Remove(name)

	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	count := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if strings.TrimSpace(sc.Text()) != "" {
			count++
		}
	}
	fmt.Printf("lines: %d\n", count)
}
