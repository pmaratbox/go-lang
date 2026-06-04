package main

import "fmt"

type level int

const (
	info level = iota
	warn
	err
)

var names = map[level]string{info: "INFO", warn: "WARN", err: "ERROR"}

func main() {
	threshold := warn
	msgs := []struct {
		lvl level
		msg string
	}{
		{info, "i"},
		{warn, "w"},
		{err, "e"},
	}
	for _, m := range msgs {
		if m.lvl >= threshold {
			fmt.Printf("%s: %s\n", names[m.lvl], m.msg)
		}
	}
}
