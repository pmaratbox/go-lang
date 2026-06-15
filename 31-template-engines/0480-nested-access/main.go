package main

import (
	"os"
	"text/template"
)

func main() {
	data := map[string]any{
		"user": map[string]any{"name": "alice"},
	}
	t := template.Must(template.New("nested").Parse("{{.user.name}}"))
	t.Execute(os.Stdout, data)
	os.Stdout.WriteString("\n")
}
