# 0588 — Nested elements with a descendant selector

Parse the fixed HTML document with the `github.com/PuerkitoBio/goquery`
library, which provides jQuery-style CSS selectors over Go's `net/html`
parser. The descendant selector `.content p` matches every `<p>` nested
inside the element with class `content`. Each paragraph's text is collected
with `Each` and joined with commas via `strings.Join` to print
`first,second`.

## Run

    go run .
