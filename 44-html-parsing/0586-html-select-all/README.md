# 0586 — Select all matching

Parse the fixed HTML document with the `github.com/PuerkitoBio/goquery`
library, which builds a queryable DOM from `goquery.NewDocumentFromReader`.
The CSS class selector `.item` matches every `<li class="item">`; `.Each`
iterates the matched selection, `.Text()` extracts each element's text, and
`strings.Join` joins them with commas to print `apple,banana,cherry`.

## Run

    go run .
