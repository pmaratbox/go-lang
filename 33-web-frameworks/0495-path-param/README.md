# 0495 — Path parameter

Capture a path parameter with Go's real `net/http` web framework. The `ServeMux` route pattern `GET /users/{id}` declares a wildcard segment that the framework binds for each matching request; the handler reads it back with `r.PathValue("id")` and echoes it. The route is exercised in-process — no port is bound — using `net/http/httptest`: `httptest.NewRequest` builds the `GET /users/42` request, `httptest.NewRecorder` captures the response in memory, and `mux.ServeHTTP` dispatches it. Printing `rec.Body.String()` yields the captured id `42`.

## Run

    go run .
