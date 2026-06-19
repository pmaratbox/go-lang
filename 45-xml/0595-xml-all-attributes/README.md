# 0595 — All attributes

Uses Go's standard-library `encoding/xml` package to parse a fixed catalog document. Struct tags map each `<book>` element and its `id` attribute (`xml:"id,attr"`); after `xml.Unmarshal`, we collect the `id` attribute of every `<book>` in document order and join them with commas.

## Run

    go run .
