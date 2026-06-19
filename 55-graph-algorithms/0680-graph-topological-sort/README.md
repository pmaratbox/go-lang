# 0680 — Topological sort

This lesson uses Go's `gonum.org/v1/gonum/graph/simple` package to build the fixed DAG (nodes a..e with directed edges a->b, b->c, a->c, c->d, d->e) and then calls `gonum.org/v1/gonum/graph/topo`'s `topo.Sort` to compute its topological order. The DAG admits a unique ordering, so the comma-joined result is `a,b,c,d,e`.

## Run

    go run .
