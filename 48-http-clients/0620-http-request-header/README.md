# 0620 — Send a request header

Using the `net/http` client against an in-process `net/http/httptest.NewServer` (a real server bound to a loopback ephemeral port), an `http.NewRequest` for `GET /token` sets a custom `X-Token: secret` header and is sent via `http.Client.Do`. The handler reads `r.Header.Get("X-Token")` and echoes it back, so the client prints the body: `secret`.

## Run

    go run .
