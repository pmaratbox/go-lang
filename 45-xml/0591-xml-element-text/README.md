# 0591 — Element text

Uses Go's standard-library `encoding/xml` to parse a fixed catalog document into structs (with `xml` field tags) via `xml.Unmarshal`, then prints the text content of the first book's `<title>` element.

## Run

    go run .
