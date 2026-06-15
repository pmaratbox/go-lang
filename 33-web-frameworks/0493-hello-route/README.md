# 0493 — Hello route

Define a `GET /` route on Go's standard `net/http` `ServeMux` that writes the text `hello`. The route is exercised entirely in-process using `net/http/httptest`: an `httptest.NewRequest` and `httptest.NewRecorder` are passed to `mux.ServeHTTP`, so no port is ever bound. The program then prints the captured response body — `hello`.

## Run

    go run .
