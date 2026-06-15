# 0565 — Base64 encode

This lesson uses Go's standard-library `encoding/base64` package to Base64-encode
data. We take the UTF-8 bytes of the string `"hello"` and encode them with the
standard alphabet (`base64.StdEncoding`), producing the padded Base64 string.

## Run

    go run .
