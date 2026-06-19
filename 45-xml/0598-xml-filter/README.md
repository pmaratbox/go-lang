# 0598 — Filter by attribute

Uses Go's standard `encoding/xml` library to unmarshal a fixed catalog document into structs (with `xml:"...,attr"` and element tags), then filters the `<book>` elements to keep only those whose `lang` attribute equals `en`, extracts each kept book's `<title>` text, and joins them with commas.

## Run

    go run .
