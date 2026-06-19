# 0679 — Connectivity

Uses Go's `gonum.org/v1/gonum/graph/simple` library to build the fixed weighted
undirected graph G (nodes a–e), then runs Dijkstra's single-source shortest-path
algorithm via `gonum.org/v1/gonum/graph/path`. Node `e` is connected to `a` when
a route exists from `a` to `e`; `pt.To(id["e"])` returns a non-empty path, so the
program prints `true`.

## Run

    go run .
