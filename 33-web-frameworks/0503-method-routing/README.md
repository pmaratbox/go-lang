# 0503 — Method routing

Route the same path to different handlers by HTTP method using Go's real `net/http` `ServeMux` with the Go 1.22 method-based pattern syntax (`GET /item`, `POST /item`). The mux is exercised in-process with `net/http/httptest`: a `httptest.NewRecorder` captures the response while `httptest.NewRequest("POST", "/item", nil)` builds the request, and `mux.ServeHTTP(rec, req)` dispatches it without binding any port. The `POST /item` pattern matches, its handler writes `post`, and we print the recorded body.

## Run

    go run .
