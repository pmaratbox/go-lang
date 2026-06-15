package main

import (
	"os"
	"text/template"
)

func main() {
	// Fixed template: render .Name, falling back to a default when it is missing.
	src := `{{if .Name}}{{.Name}}{{else}}anonymous{{end}}`

	t := template.Must(template.New("default-value").Parse(src))

	// Fixed data with NO name field.
	data := map[string]string{}

	t.Execute(os.Stdout, data)
	os.Stdout.WriteString("\n")
}
