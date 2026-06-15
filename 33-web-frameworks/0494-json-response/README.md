# 0494 — JSON response

Serve a JSON body with Go's real `net/http` framework. An `http.ServeMux` registers `GET /user`, whose handler sets the `application/json` content type and writes the compact object `{"name":"alice"}`. The route is exercised entirely in-process using `net/http/httptest`: `httptest.NewRecorder()` captures the response while `mux.ServeHTTP` is invoked with `httptest.NewRequest`, so no fixed port is bound. The program prints the recorded response body straight from the framework.

## Run

    go run .
