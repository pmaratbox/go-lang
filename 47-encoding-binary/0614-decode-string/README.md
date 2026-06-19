# 0614 — Decode bytes

Uses the `github.com/vmihailenco/msgpack/v5` MessagePack library to DECODE bytes back into a value. The fixed hex string `a568656c6c6f` is converted to bytes with `hex.DecodeString`, then `msgpack.Unmarshal` decodes the `fixstr` (marker `a5` + UTF-8 bytes) into a Go `string`, which prints as `hello`.

## Run

    go run .
