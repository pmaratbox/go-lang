# 0612 — Encode a map

Uses the `github.com/vmihailenco/msgpack/v5` MessagePack library to encode the single-key map `{"a": 1}` to its binary MessagePack representation, then prints the lowercase hex of those bytes. The encoding is a fixmap (`81`) holding one key/value pair: the string key `a` (`a161`) and the integer value `1` (`01`).

## Run

    go run .
