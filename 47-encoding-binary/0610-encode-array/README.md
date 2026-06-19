# 0610 — Encode an array

Using the `github.com/vmihailenco/msgpack/v5` MessagePack library, `msgpack.Marshal` encodes the array `[1, 2, 3]` to its binary form, which `hex.EncodeToString` renders as lowercase hex: `93010203` (a fixarray header `93` for a 3-element array, followed by the positive fixints `01 02 03`).

## Run

    go run .
