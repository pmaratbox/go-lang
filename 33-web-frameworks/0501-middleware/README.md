# 0501 — Middleware

Apply middleware that transforms the response using Go's standard `net/http` framework. The `prefixMiddleware` wraps the `http.ServeMux` handler: it captures whatever the inner handler writes into an `httptest.ResponseRecorder`, then relays it to the outer `ResponseWriter` prefixed with `[mw] `. The `GET /` route handler writes only `hello`, so after the middleware the body becomes `[mw] hello`. The route is exercised in-process with `httptest.NewRecorder()` + `httptest.NewRequest()` and `handler.ServeHTTP` — no fixed listening port — and the printed value comes from the real recorded HTTP response body, not a literal.

## Run

    go run .
