# 0624 — PUT request

Using Go's standard `net/http` client, `http.NewRequest("PUT", ...)` plus `http.Client.Do` sends a PUT request to an in-process `net/http/httptest.NewServer` (a real loopback server on an ephemeral port). The `/item` route replies with the text `updated` for PUT, and the client reads the response body and prints it: `updated`.

## Run

    go run .
