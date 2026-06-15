# 0465 — Round trip

Uses the standard library `encoding/json` package to round-trip a value:
`json.Marshal` serializes a `Person{Age: 30, Name: "alice"}` struct (fields
declared alphabetically with `json:"..."` tags, yielding compact keys in
order) into a JSON byte slice, and `json.Unmarshal` parses those bytes back
into a fresh `Person`. The recovered name is printed to confirm both
directions worked.

## Run

    go run .
