# 0621 — Read a response header

Uses Go's standard `net/http` client against an in-process `net/http/httptest.NewServer` (a real loopback server on an ephemeral port, no external network). The `/info` route sets the custom response header `X-Count: 7`. The client performs `http.Get` and reads the header value with `resp.Header.Get("X-Count")`, printing `7`. A custom header is used instead of `Content-Type` to avoid any charset munging.

## Run

    go run .
