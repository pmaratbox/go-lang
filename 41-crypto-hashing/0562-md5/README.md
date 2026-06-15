# 0562 — MD5

Compute an MD5 digest with Go's standard-library `crypto/md5` package.
`md5.Sum` hashes the UTF-8 bytes of `"hello"` into a fixed 16-byte array,
and `encoding/hex.EncodeToString` renders it as a lowercase hex string (no
colons or spaces). The digest is deterministic for this fixed input.

## Run

    go run .
