package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	// Go's stdlib net/http ServeMux supports typed path patterns since 1.22.
	// The {id} wildcard is captured and read back with r.PathValue.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.PathValue("id"))
	})

	// Exercise the route IN-PROCESS: no port is bound. httptest builds a
	// request and an in-memory ResponseRecorder, and we drive the mux directly.
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("GET", "/users/42", nil))

	// Print the body produced by the actual HTTP response.
	fmt.Println(rec.Body.String())
}
