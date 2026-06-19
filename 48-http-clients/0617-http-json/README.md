# 0617 — Parse JSON response

Using the stdlib `net/http` client against an in-process `net/http/httptest.NewServer` (a real server on a loopback ephemeral port, no external network), this lesson issues `http.Get` to the `/user` route that returns `{"name":"Alice","age":30}`, decodes the JSON body with `encoding/json`, and prints the `name` field: `Alice`.

## Run

    go run .
