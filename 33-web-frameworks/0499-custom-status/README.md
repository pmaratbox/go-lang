# 0499 — Custom status

Register a `POST /create` route on Go's standard `net/http` `ServeMux` whose handler calls `w.WriteHeader(http.StatusCreated)` to return a custom `201 Created` status. The route is exercised entirely in-process with `net/http/httptest`: an `httptest.NewRequest` and `httptest.NewRecorder` are passed to `mux.ServeHTTP`, so no port is ever bound. The program then prints the captured response status code from `rec.Code` — `201`.

## Run

    go run .
