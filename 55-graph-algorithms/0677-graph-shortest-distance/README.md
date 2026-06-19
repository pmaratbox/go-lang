# 0677 — Shortest distance

Uses Go's `gonum.org/v1/gonum/graph/simple` library to build the fixed weighted
undirected graph G (nodes a–e) and `gonum.org/v1/gonum/graph/path`'s
`DijkstraFrom` to compute the single-source shortest paths from `a`. Calling
`pt.To(id["e"])` returns the weighted distance to `e`; the unique shortest path
is `a-b-c-d-e` with total cost `4`.

## Run

    go run .
