# 0498 — 404 status

Build an `http.ServeMux` with Go's standard `net/http` framework that registers only a `/hello` route, then exercise it in-process with `net/http/httptest` (`httptest.NewRecorder` + `httptest.NewRequest`, dispatched through `mux.ServeHTTP` — no port is ever bound). A request to the undefined route `GET /missing` finds no matching pattern, so the mux's built-in handling responds with status `404`, and the program prints the recorder's real `Code` -> `404`.

## Run

    go run .
