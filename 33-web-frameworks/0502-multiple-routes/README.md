# 0502 — Multiple routes

Register two routes on Go's standard `net/http` `ServeMux` — `GET /` writing `home` and `GET /about` writing `about` — then request both in order entirely in-process using `net/http/httptest`. Each call pairs an `httptest.NewRequest` with an `httptest.NewRecorder` and dispatches via `mux.ServeHTTP`, so no port is ever bound. The captured response body of each request is printed on its own line, yielding `home` then `about`.

## Run

    go run .
