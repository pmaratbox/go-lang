# 0596 — Count elements

Parse the fixed catalog document with Go's standard-library `encoding/xml`
package. `xml.Unmarshal` decodes the document into a `Catalog` struct whose
`Books []Book` field (tagged `xml:"book"`) collects every `<book>` element.
Counting the slice with `len(c.Books)` yields the number of books, printing `2`.

## Run

    go run .
