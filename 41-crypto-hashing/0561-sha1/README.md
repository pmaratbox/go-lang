# 0561 — SHA-1

Compute the SHA-1 hash of `"hello"` using Go's standard-library `crypto/sha1` package and print the lowercase hex digest. `sha1.Sum` returns a fixed `[20]byte` array, and `encoding/hex`'s `EncodeToString` formats it as lowercase hex with no colons or spaces. SHA-1 is a legacy algorithm and is cryptographically broken for collision resistance, but it is still useful for non-security checksums.

## Run

    go run .
