// Go — net/http client + net/http/httptest in-process server (loopback ephemeral port). Run: go run .
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	// The server only defines /hello, so any other route returns 404.
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/missing")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}
