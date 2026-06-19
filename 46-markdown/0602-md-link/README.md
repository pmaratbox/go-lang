# 0602 — Link

Uses the `github.com/yuin/goldmark` Markdown library to render the inline link `[text](http://x.com)` to HTML via `goldmark.Convert`, then strips the trailing newline the renderer appends and prints the resulting `<a>` element.

## Run

    go run .
