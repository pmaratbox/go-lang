// Go — net/http client + net/http/httptest in-process server (loopback ephemeral port). Run: go run .
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		w.Write(b)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	resp, _ := http.Post(srv.URL+"/echo", "text/plain", strings.NewReader("ping"))
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
