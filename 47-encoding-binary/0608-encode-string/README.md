# 0608 — Encode a string

Uses the `github.com/vmihailenco/msgpack/v5` MessagePack library to encode the fixed string `"hello"` via `msgpack.Marshal`, then prints the lowercase hex of the resulting bytes with `hex.EncodeToString`. MessagePack encodes a short string as a `fixstr`: the marker `a5` (string of length 5) followed by the UTF-8 bytes `68656c6c6f`.

## Run

    go run .
