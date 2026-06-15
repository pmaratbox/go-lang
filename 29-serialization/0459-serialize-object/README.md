# 0459 — Serialize an object

Serialize a typed `Person{age, name}` value into a compact JSON string using Go's standard `encoding/json` package. `json.Marshal` emits compact JSON (no spaces); declaring the struct fields alphabetically with `json:"..."` tags keeps the output keys in alphabetical order, producing `{"age":30,"name":"alice"}`.

## Run

    go run .
