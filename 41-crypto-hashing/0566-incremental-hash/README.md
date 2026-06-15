# 0566 — Incremental hashing

Compute a SHA-256 digest incrementally with Go's standard-library
`crypto/sha256` package. Instead of the one-shot `sha256.Sum256`, this creates
a streaming `hash.Hash` via `sha256.New()` and feeds it the data in two
separate `Write` calls (`"foo"` then `"bar"`). `Sum(nil)` finalizes the digest,
and `encoding/hex.EncodeToString` renders it as a lowercase hex string (no
colons or spaces). The result equals the SHA-256 of `"foobar"` and is
deterministic for this fixed input.

## Run

    go run .
