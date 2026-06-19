# 0676 — Edge existence

This lesson uses Go's `gonum.org/v1/gonum/graph/simple` package to build the fixed weighted undirected graph G (nodes a..e with six edges). It calls the graph's own `HasEdgeBetween` method to test edge existence: edge b-c exists while edge a-e does not, printing `true false`.

## Run

    go run .
