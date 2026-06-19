# 0584 — Select by id

Parse the fixed HTML document with the `github.com/PuerkitoBio/goquery`
library, which builds a queryable DOM from `goquery.NewDocumentFromReader`.
The CSS id selector `#status` matches the single element whose `id` is
`status`, and `.Text()` returns its text content `active`.

## Run

    go run .
