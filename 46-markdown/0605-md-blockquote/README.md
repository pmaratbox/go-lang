# 0605 — Blockquote

Uses the `github.com/yuin/goldmark` Markdown library to render the blockquote `> quote` to HTML via `goldmark.Convert`, then strips the trailing newline the renderer appends and prints the resulting `<blockquote>` element wrapping a paragraph.

## Run

    go run .
