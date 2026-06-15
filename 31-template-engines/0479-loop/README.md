# 0479 — Loop

Iterate a list in a template using Go's real `text/template` engine. The template uses the `{{range $i, $n := .Nums}}` action to walk the fixed slice `[1, 2, 3]`, printing each element on its own line; a `{{if $i}}\n{{end}}` guard emits the newline separator before every element except the first so the rendered output is exactly `1\n2\n3`.

## Run

    go run .
