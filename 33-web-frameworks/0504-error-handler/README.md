# 0504 — Error handler

Register a `GET /boom` route on Go's standard `net/http` `ServeMux` whose handler `panic`s with a real `error`, and wrap the mux in a `recoverMiddleware` that acts as the framework's error handler: it `recover`s the thrown error and converts it into a `500 Internal Server Error` via `http.Error`. The route is exercised entirely in-process with `net/http/httptest` — an `httptest.NewRequest` and `httptest.NewRecorder` are passed to `handler.ServeHTTP`, so no port is ever bound. The program prints the captured response status code from `rec.Code` — `500` — produced by the error handling, not hardcoded.

## Run

    go run .
