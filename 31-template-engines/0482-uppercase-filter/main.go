package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	src := "{{upper .Name}}"
	t := template.Must(template.New("x").Funcs(template.FuncMap{
		"upper": strings.ToUpper,
	}).Parse(src))
	t.Execute(os.Stdout, map[string]string{"Name": "alice"})
	os.Stdout.WriteString("\n")
}
