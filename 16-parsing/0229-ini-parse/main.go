package main

import (
	"fmt"
	"strings"
)

func main() {
	ini := "[s]\nk=v"
	section := ""
	for _, line := range strings.Split(ini, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			section = line[1 : len(line)-1]
			continue
		}
		if k, v, ok := strings.Cut(line, "="); ok {
			fmt.Println(section + "." + k + "=" + v)
		}
	}
}
