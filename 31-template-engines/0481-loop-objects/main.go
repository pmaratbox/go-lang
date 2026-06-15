package main

import (
	"os"
	"text/template"
)

type User struct {
	Name string
	Age  int
}

func main() {
	src := "{{range $i, $u := .Users}}{{if $i}}\n{{end}}{{$u.Name}}: {{$u.Age}}{{end}}"
	t := template.Must(template.New("loop-objects").Parse(src))
	data := map[string]any{
		"Users": []User{
			{Name: "alice", Age: 30},
			{Name: "bob", Age: 25},
		},
	}
	t.Execute(os.Stdout, data)
	os.Stdout.WriteString("\n")
}
