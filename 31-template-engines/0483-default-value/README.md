# 0483 — Default value

Uses Go's standard-library `text/template` engine to render a fixed template
against fixed data that contains no `Name` field. The template relies on the
engine's `{{if .Name}}...{{else}}...{{end}}` conditional: because `.Name` is
absent (its zero value is falsy), the `else` branch supplies the default,
printing `anonymous`.

## Run

    go run .
