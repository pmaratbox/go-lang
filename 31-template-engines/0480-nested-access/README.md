# 0480 — Nested access

Render a fixed template with Go's real `text/template` engine, using its field-access syntax to reach a nested value. The template `{{.user.name}}` walks from the root data map into the `user` map and reads its `name` key, so the chained dot path resolves to the inner string. The data is fixed as `{user: {name: alice}}`, so the program always prints `alice`.

## Run

    go run .
