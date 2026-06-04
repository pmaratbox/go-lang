# 0327 — FlatMap

FlatMap [1,2,3] with x -> [x, x*10] and print the flattened result `1 10 2 20 3 30`. Go's variadic `append(out, f(x)...)` flattens each produced slice in place.

## Run

    go run .
