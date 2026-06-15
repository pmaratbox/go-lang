# 0460 — Deserialize an object

Parse a JSON string into a typed object using Go's standard `encoding/json` library. `json.Unmarshal` reads `{"age":30,"name":"alice"}` into a `Person` struct (fields mapped via `json:"..."` tags), then we print `name age` -> `alice 30`.

## Run

    go run .
