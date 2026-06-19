# 0607 — Encode an integer

Using the `github.com/vmihailenco/msgpack/v5` MessagePack library, `msgpack.Marshal` encodes the integer `42` to its binary form, which `hex.EncodeToString` renders as lowercase hex: `2a` (a positive fixint that stores small non-negative integers in a single byte).

## Run

    go run .
