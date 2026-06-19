# 0593 — Nested element text

Uses Go's standard-library `encoding/xml` to parse a fixed catalog document into structs (with `xml` field tags) via `xml.Unmarshal`, then prints the text of the first book's nested `<price>` element as an integer.

## Run

    go run .
