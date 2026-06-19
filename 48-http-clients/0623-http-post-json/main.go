// Go — net/http client + net/http/httptest in-process server (loopback ephemeral port). Run: go run .
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/double", func(w http.ResponseWriter, r *http.Request) {
		var m map[string]int
		json.NewDecoder(r.Body).Decode(&m)
		json.NewEncoder(w).Encode(map[string]int{"doubled": m["x"] * 2})
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	resp, _ := http.Post(srv.URL+"/double", "application/json", strings.NewReader(`{"x":5}`))
	var dm map[string]int
	json.NewDecoder(resp.Body).Decode(&dm)
	resp.Body.Close()
	fmt.Println(dm["doubled"])
}
