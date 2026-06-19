# 0622 — Handle 404

Using the `net/http` client against an in-process `net/http/httptest.NewServer` (a real server bound to a loopback ephemeral port), `http.Get` calls `GET /missing`. The server's `ServeMux` only registers `/hello`, so the unknown route is answered with `404 Not Found`. The client reports the integer status code: `404`.

## Run

    go run .
