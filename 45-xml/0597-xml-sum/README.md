# 0597 — Sum numeric children

Uses Go's standard `encoding/xml` library to parse a fixed catalog document into structs (`xml.Unmarshal`). Each `<book>`'s nested `<price>` element is mapped to an `int` field via the `xml:"price"` tag, and the prices are summed (30 + 45 = 75).

## Run

    go run .
