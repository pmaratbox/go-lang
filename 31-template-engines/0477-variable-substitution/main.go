package main

import (
	"os"
	"text/template"
)

func main() {
	const src = "Hello {{.Name}}"
	t := template.Must(template.New("greeting").Parse(src))
	t.Execute(os.Stdout, map[string]string{"Name": "alice"})
	os.Stdout.WriteString("\n")
}
