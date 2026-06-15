# 0500 — Request header

Read a request header with Go's real `net/http` framework, exercised in-process via the `net/http/httptest` test recorder instead of a bound port. An `http.ServeMux` registers `GET /whoami`, whose handler echoes the incoming `X-Name` header by calling `r.Header.Get("X-Name")`. The test request is built with `httptest.NewRequest`, its `X-Name` header is set to `alice`, and `mux.ServeHTTP` runs the handler against an `httptest.NewRecorder()`. Printing `rec.Body.String()` yields the echoed header value `alice`.

## Run

    go run .
