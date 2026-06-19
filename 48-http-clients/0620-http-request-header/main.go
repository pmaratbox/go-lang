// Go — net/http client + net/http/httptest in-process server (loopback ephemeral port). Run: go run .
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	mux := http.NewServeMux()
	// GET /token echoes the request header X-Token in the body.
	mux.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Header.Get("X-Token"))
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	c := &http.Client{}
	req, _ := http.NewRequest("GET", srv.URL+"/token", nil)
	req.Header.Set("X-Token", "secret")
	resp, _ := c.Do(req)
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))
}
