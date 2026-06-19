# 0589 — Extract table cells

Parse the fixed TABLE document with the `github.com/PuerkitoBio/goquery`
library, which builds a queryable DOM from `goquery.NewDocumentFromReader`.
The CSS tag selector `td` matches every table cell; iterating with `.Each`
collects each cell's `.Text()` in row-major order, and joining with commas
yields `r1c1,r1c2,r2c1,r2c2`.

## Run

    go run .
