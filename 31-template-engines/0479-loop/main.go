package main

import (
	"os"
	"text/template"
)

func main() {
	src := "{{range $i, $n := .Nums}}{{if $i}}\n{{end}}{{$n}}{{end}}"
	t := template.Must(template.New("loop").Parse(src))
	data := map[string][]int{"Nums": {1, 2, 3}}
	if err := t.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
	os.Stdout.WriteString("\n")
}
