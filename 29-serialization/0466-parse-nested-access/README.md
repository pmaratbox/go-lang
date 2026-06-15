# 0466 — Parse Nested Access

Parse `{"user":{"name":"alice","roles":["admin","user"]}}` with the standard library's `encoding/json`, decoding into a generic `map[string]any` tree (the dynamic API). Navigate the tree with type assertions to reach `user.name` and the first role, printing `alice admin`.

## Run

    go run .
