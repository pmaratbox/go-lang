# 0609 — Encode a boolean

Using the `github.com/vmihailenco/msgpack/v5` MessagePack library, `msgpack.Marshal` encodes the boolean `true` to its binary form, which `hex.EncodeToString` renders as lowercase hex: `c3` (MessagePack uses the single byte `c3` for `true` and `c2` for `false`).

## Run

    go run .
