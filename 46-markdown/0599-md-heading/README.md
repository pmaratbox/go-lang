# 0599 — Heading

Uses the `github.com/yuin/goldmark` Markdown library to render the fixed Markdown input `# Hello` to HTML, strips the trailing newline that the renderer appends, and prints the result. The construct rendered here is an ATX heading (`#`), which goldmark converts to an `<h1>` element.

## Run

    go run .
