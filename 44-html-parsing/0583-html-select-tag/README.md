# 0583 — Select by tag

Parse the fixed HTML document with the `github.com/PuerkitoBio/goquery`
library, which builds a queryable DOM from `goquery.NewDocumentFromReader`.
The CSS tag selector `h1` matches the heading element, and `.Text()` returns
its text content `Hello`.

## Run

    go run .
