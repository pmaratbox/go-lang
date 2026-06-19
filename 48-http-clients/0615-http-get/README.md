# 0615 — GET request

Uses Go's standard-library `net/http` client (`http.Get`) to make a GET request against an in-process server started with `net/http/httptest.NewServer`, which listens on a loopback ephemeral port. The server exposes a single route, `GET /hello`, that returns the fixed text body `hello world`; the client reads the response body and prints it.

## Run

    go run .
