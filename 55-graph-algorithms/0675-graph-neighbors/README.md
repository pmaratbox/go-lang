# 0675 — Neighbors

Uses Go's `gonum.org/v1/gonum/graph/simple` library to build the fixed weighted
undirected graph G (nodes a–e). Calling `g.From(id["a"])` returns the adjacency
iterator for node `a`; the resulting node IDs are mapped back to labels, sorted
for determinism, and comma-joined to print `b,c`.

## Run

    go run .
