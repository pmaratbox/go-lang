// Go — net/http client + net/http/httptest in-process server (loopback ephemeral port). Run: go run .
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"name": "Alice", "age": 30})
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	resp, _ := http.Get(srv.URL + "/user")
	defer resp.Body.Close()
	var user map[string]any
	json.NewDecoder(resp.Body).Decode(&user)
	fmt.Println(user["name"])
}
