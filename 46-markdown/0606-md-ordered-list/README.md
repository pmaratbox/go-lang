# 0606 — Ordered list

Uses the `github.com/yuin/goldmark` Markdown library to render the fixed Markdown input `1. a\n2. b` to HTML, strips the trailing newline that the renderer appends, and prints the result. The construct rendered here is an ordered (numbered) list, which goldmark converts to an `<ol>` element wrapping `<li>` items.

## Run

    go run .
