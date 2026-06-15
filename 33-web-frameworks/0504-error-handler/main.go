package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

// recoverMiddleware is the framework's error handler: it wraps the next
// handler and, if that handler panics with a real error, recovers and turns
// the thrown error into an HTTP 500 response.
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Silence the default logger so a recovered panic's stack/log noise
	// never leaks into stdout.
	log.SetOutput(io.Discard)

	mux := http.NewServeMux()
	mux.HandleFunc("/boom", func(w http.ResponseWriter, r *http.Request) {
		// The handler throws a real error.
		panic(fmt.Errorf("boom"))
	})

	handler := recoverMiddleware(mux)

	// Exercise the route in-process: no port is ever bound.
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/boom", nil)
	handler.ServeHTTP(rec, req)

	// Print the real status code produced by the framework's error handling.
	fmt.Println(rec.Code)
}
