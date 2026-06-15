# 0496 — Query parameter

Reads a query-string parameter using Go's real `net/http` framework. The route `GET /greet` is registered on an `http.ServeMux` and its handler returns `hello ` plus the `name` query value via `r.URL.Query().Get("name")`. The route is exercised in-process with `net/http/httptest` (`httptest.NewRequest` + `httptest.NewRecorder` driven through `mux.ServeHTTP`), so no fixed port is bound; the program prints the recorded response body — here `hello alice`.

## Run

    go run .
