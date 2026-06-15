# 0477 — Variable substitution

Uses Go's standard-library `text/template` engine to render a fixed template
containing a single `{{.Name}}` action against fixed data `{Name: alice}`. The
template source is parsed with `template.Parse` and executed with
`t.Execute`, which substitutes the variable into the surrounding text to
produce `Hello alice`.

## Run

    go run .
