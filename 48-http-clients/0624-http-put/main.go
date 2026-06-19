// Go — net/http client + net/http/httptest in-process server (loopback ephemeral port).
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			fmt.Fprint(w, "updated")
		}
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	c := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, srv.URL+"/item", nil)
	resp, _ := c.Do(req)
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(body)) // updated
}
