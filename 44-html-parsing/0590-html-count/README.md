# 0590 — Count matches

Parse the fixed HTML document with the `github.com/PuerkitoBio/goquery`
library, which builds a queryable DOM from `goquery.NewDocumentFromReader`.
The CSS class selector `.item` matches the three list items, and `.Length()`
returns the number of matched elements, printing `3`.

## Run

    go run .
