# 0564 — PBKDF2

Derive a key with PBKDF2 using Go's standard-library `crypto/pbkdf2` package
(available since Go 1.24). `pbkdf2.Key` runs PBKDF2-HMAC-SHA256 over the
password `"password"` and salt `"salt"` for 1000 iterations to produce a 32-byte
derived key, and `encoding/hex.EncodeToString` renders it as a lowercase hex
string (no colons or spaces). The result is deterministic for this fixed input.

## Run

    go run .
