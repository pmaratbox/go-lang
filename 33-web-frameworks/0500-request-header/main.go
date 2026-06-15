package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/whoami", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Header.Get("X-Name"))
	})

	req := httptest.NewRequest("GET", "/whoami", nil)
	req.Header.Set("X-Name", "alice")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	fmt.Println(rec.Body.String())
}
