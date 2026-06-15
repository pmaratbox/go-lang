package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// prefixMiddleware wraps a handler, capturing its written body and prefixing
// it with "[mw] " before relaying it to the real ResponseWriter. The handler
// itself only writes "hello"; the prefix is added entirely by the middleware.
func prefixMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()
		next.ServeHTTP(rec, r)
		fmt.Fprint(w, "[mw] "+rec.Body.String())
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})

	handler := prefixMiddleware(mux)

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))

	fmt.Println(rec.Body.String())
}
