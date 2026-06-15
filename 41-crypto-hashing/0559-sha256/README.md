# 0559 — SHA-256

Compute a SHA-256 digest with Go's standard-library `crypto/sha256` package.
`sha256.Sum256` hashes the UTF-8 bytes of `"hello"` into a fixed 32-byte array,
and `encoding/hex.EncodeToString` renders it as a lowercase hex string (no
colons or spaces). The digest is deterministic for this fixed input.

## Run

    go run .
