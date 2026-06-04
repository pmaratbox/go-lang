# 0196 — Error Wrapping

Wrap an inner error "inner" inside an outer context and print the combined message `outer: inner`. Go wraps with `fmt.Errorf` using the `%w` verb to preserve the underlying error.

## Run

    go run .
