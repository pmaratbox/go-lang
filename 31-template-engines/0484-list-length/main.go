package main

import (
	"os"
	"text/template"
)

func main() {
	src := "{{len .Items}}"
	t := template.Must(template.New("list-length").Parse(src))
	data := map[string]any{"Items": []int{1, 2, 3}}
	t.Execute(os.Stdout, data)
	os.Stdout.WriteString("\n")
}
