# 0603 — Inline code

Uses the `github.com/yuin/goldmark` Markdown library to render a backtick-wrapped
inline code span (`` `code` ``) to HTML via `goldmark.Convert`, then strips the
trailing newline the renderer appends to print `<p><code>code</code></p>`.

## Run

    go run .
