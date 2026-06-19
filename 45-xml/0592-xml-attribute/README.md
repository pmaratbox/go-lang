# 0592 — Read an attribute

Uses Go's standard-library `encoding/xml` to parse a fixed catalog document into structs (with `xml` field tags, where `id,attr` binds the element attribute) via `xml.Unmarshal`, then prints the `id` attribute of the first `<book>` element.

## Run

    go run .
