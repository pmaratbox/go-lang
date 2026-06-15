# 0563 — HMAC-SHA256

Compute an HMAC-SHA256 authentication code with Go's standard-library
`crypto/hmac` and `crypto/sha256` packages. `hmac.New(sha256.New, key)` builds a
keyed-hash MAC using SHA-256 as the underlying hash and `"key"` as the secret;
writing the message `"hello"` and calling `Sum` produces a fixed 32-byte tag,
which `encoding/hex.EncodeToString` renders as a lowercase hex string (no colons
or spaces). The result is deterministic for this fixed key and message.

## Run

    go run .
