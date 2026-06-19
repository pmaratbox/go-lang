# 0604 — Unordered list

Uses the `github.com/yuin/goldmark` Markdown library to render the fixed Markdown input `- a\n- b` to HTML, strips the trailing newline that the renderer appends, and prints the result. The construct rendered here is an unordered (bullet) list, which goldmark converts to a `<ul>` element wrapping `<li>` items.

## Run

    go run .
