# 0594 — Find all elements

Parse the fixed catalog document with Go's standard-library `encoding/xml`
package. `xml.Unmarshal` maps every `<book>` element onto a slice field
tagged `xml:"book"`, with each book's `<title>` text captured via the
`xml:"title"` struct tag. Iterating the slice collects all titles in
document order, joined with commas via `strings.Join` to print `Go,Rust`.

## Run

    go run .
