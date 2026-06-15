package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /item", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "get") })
	mux.HandleFunc("POST /item", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "post") })

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("POST", "/item", nil))
	fmt.Println(rec.Body.String())
}
