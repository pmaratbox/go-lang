# 0678 — Shortest path route

Uses Go's `gonum.org/v1/gonum/graph` library to build the fixed weighted
undirected graph G (nodes a–e) and runs Dijkstra's algorithm via
`path.DijkstraFrom(a, g)`. Calling `.To(e)` on the resulting shortest-path tree
returns the unique minimum-cost route, whose node IDs are mapped back to labels
and joined with `-` to print `a-b-c-d-e` (total cost 4).

## Run

    go run .
