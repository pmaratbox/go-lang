# 0482 — Uppercase filter

Render a fixed template with Go's standard library `text/template` engine, applying an uppercase "filter" via a custom helper function. The template `{{upper .Name}}` calls a function named `upper` registered through `Funcs(template.FuncMap{"upper": strings.ToUpper})`, which transforms the data value `name=alice` into `ALICE` at render time. `template.Execute` writes the rendered result to stdout.

## Run

    go run .
