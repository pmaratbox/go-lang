# 0623 — POST JSON, parse JSON

Using the standard library `net/http` client (`http.Post`) against an in-process `net/http/httptest.NewServer` (a real server on a loopback ephemeral port, no external network), this POSTs `{"x":5}` to the `/double` route, which reads the JSON and replies `{"doubled":10}`. The client parses the JSON reply with `encoding/json` and prints the doubled value: `10`.

## Run

    go run .
