# 0497 — POST JSON body

Parses a JSON request body using Go's real `net/http` framework. The route `POST /sum` is registered on an `http.ServeMux`; its handler decodes the request body with `json.NewDecoder(r.Body).Decode(&in)` into a struct and writes back `a + b`. The route is exercised in-process with `net/http/httptest` (`httptest.NewRequest` carrying the JSON body `{"a":2,"b":3}` plus `httptest.NewRecorder`, driven through `mux.ServeHTTP`), so no fixed port is bound; the program prints the recorded response body — here `5`.

## Run

    go run .
