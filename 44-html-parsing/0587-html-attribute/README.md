# 0587 — Read an attribute

Query the fixed HTML document with `github.com/PuerkitoBio/goquery`, a
jQuery-style HTML library built on Go's `golang.org/x/net/html` parser.
`goquery.NewDocumentFromReader` parses `DOC`, the `a` CSS selector matches the
single anchor element, and `.Attr("href")` reads its `href` attribute, printing
`https://example.com`.

## Run

    go run .
