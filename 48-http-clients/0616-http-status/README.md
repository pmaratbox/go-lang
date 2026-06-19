# 0616 — Response status code

Using the `net/http` client against an in-process `net/http/httptest.NewServer` (a real server bound to a loopback ephemeral port), `http.Get` calls `GET /hello` and reads `resp.StatusCode`. The handler returns `200 OK`, so the client reports the integer status code: `200`.

## Run

    go run .
