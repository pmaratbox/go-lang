package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	// Build a real net/http ServeMux with two routes.
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "home")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "about")
	})

	// Exercise each route in-process with httptest (no port is ever bound).
	for _, path := range []string{"/", "/about"} {
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest("GET", path, nil))
		fmt.Println(rec.Body.String())
	}
}
