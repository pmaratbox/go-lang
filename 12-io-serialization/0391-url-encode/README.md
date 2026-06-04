# 0391 — URL Percent-Encode

Percent-encode the string "a b&c" to `a%20b%26c`. Go's `%02X` verb yields uppercase hex for each reserved byte while unreserved characters pass through unchanged.

## Run

    go run .
