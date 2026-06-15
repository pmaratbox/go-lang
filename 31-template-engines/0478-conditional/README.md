# 0478 — Conditional

Render a fixed template with Go's real `text/template` engine, using its `{{if}}/{{else}}/{{end}}` conditional action. The template branches on the boolean field `.LoggedIn`; given the fixed data `{logged_in: true}` it takes the `if` branch and prints `welcome` (otherwise it would print `guest`). The engine parses the template string and executes it against the data, writing the result to stdout.

## Run

    go run .
