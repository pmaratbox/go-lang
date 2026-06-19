# 0674 — Node degree

This lesson uses Go's `gonum.org/v1/gonum/graph/simple` library to build the fixed weighted undirected graph G and read the degree of a node. The undirected graph's `From(id)` iterator yields the neighbours of a node, so `From(b).Len()` returns the degree of `b` directly from the graph structure. Node `b` connects to `a`, `c`, and `d`, so its degree is `3`.

## Run

    go run .
