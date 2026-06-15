package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

type sumReq struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /sum", func(w http.ResponseWriter, r *http.Request) {
		var in sumReq
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "%d", in.A+in.B)
	})

	body := strings.NewReader(`{"a":2,"b":3}`)
	req := httptest.NewRequest("POST", "/sum", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	fmt.Println(rec.Body.String())
}
