// Go — net/http client + net/http/httptest in-process server (loopback ephemeral port).
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	// In-process server exposing only the route this lesson needs.
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	// Real HTTP client call — the body comes from the response, not a literal.
	resp, err := http.Get(srv.URL + "/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
