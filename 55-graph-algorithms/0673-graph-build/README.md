# 0673 — Build a graph

This lesson uses Go's `gonum.org/v1/gonum/graph/simple` package to build the fixed weighted undirected graph G (nodes a..e with six edges). It then queries the graph's own `Nodes().Len()` and `Edges().Len()` accessors to report the node and edge counts, printing `5 6`.

## Run

    go run .
