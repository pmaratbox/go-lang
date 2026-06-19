# 0618 — POST a body

Uses Go's standard-library `net/http` client (`http.Post`) to send a request body to an in-process server started with `net/http/httptest.NewServer`, which listens on a loopback ephemeral port. The server exposes a single route, `POST /echo`, that returns the request body verbatim; the client POSTs the text `ping` and prints the echoed response body.

## Run

    go run .
