package main

import (
	"os"
	"text/template"
)

func main() {
	src := "{{if .LoggedIn}}welcome{{else}}guest{{end}}"
	t := template.Must(template.New("conditional").Parse(src))
	t.Execute(os.Stdout, map[string]bool{"LoggedIn": true})
	os.Stdout.WriteString("\n")
}
