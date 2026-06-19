# 0613 — Encode null

Using the `github.com/vmihailenco/msgpack/v5` MessagePack library, `msgpack.Marshal` encodes the nil value to its binary form, which `hex.EncodeToString` renders as lowercase hex: `c0` (the single-byte nil format).

## Run

    go run .
