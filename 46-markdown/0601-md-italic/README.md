# 0601 — Italic

Uses the `github.com/yuin/goldmark` Markdown library to render the fixed input `*italic*` to HTML via `goldmark.Convert`, then strips the trailing newline goldmark appends and prints the emphasis (`<em>`) output.

## Run

    go run .
