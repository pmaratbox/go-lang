# 0619 — Query parameters

Uses Go's standard-library `net/http` client (`http.Get`) to call an in-process server started with `net/http/httptest.NewServer`, which listens on a loopback ephemeral port. The server exposes a single route, `GET /greet`, that reads the `name` query parameter via `r.URL.Query().Get("name")` and returns `hi <name>`. The client requests `/greet?name=Bob` and prints the response body.

## Run

    go run .
