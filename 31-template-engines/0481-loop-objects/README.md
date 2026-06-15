# 0481 — Loop objects

Render a fixed template over a fixed slice of user structs using Go's real template engine, the standard library `text/template`. The template uses the `{{range $i, $u := .Users}}` action to iterate the list, emitting `name: age` for each user; a `{{if $i}}\n{{end}}` guard inserts a newline before every item except the first so the lines are separated without a trailing blank line.

## Run

    go run .
