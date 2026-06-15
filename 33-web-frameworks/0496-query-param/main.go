package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, "hello "+name)
	})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/greet?name=alice", nil)
	mux.ServeHTTP(rec, req)
	fmt.Println(rec.Body.String())
}
