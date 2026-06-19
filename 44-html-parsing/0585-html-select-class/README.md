# 0585 — Select by class

Query a fixed HTML document with the `github.com/PuerkitoBio/goquery`
library, which parses HTML and supports jQuery-style CSS selectors.
`d.Find(".item")` matches every element with class `item`; `.First()`
narrows the selection to the first match and `.Text()` returns its text,
printing `apple`.

## Run

    go run .
