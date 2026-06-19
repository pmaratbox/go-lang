# 0600 — Bold

Uses the `github.com/yuin/goldmark` Markdown library to render the fixed input `**bold**` to HTML via `goldmark.Convert`, then strips the trailing newline the renderer appends. The `**...**` construct produces a `<strong>` (strong emphasis) element wrapped in a paragraph.

## Run

    go run .
