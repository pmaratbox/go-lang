# 0306 — Reconstruct Shortest Path

On the Dijkstra graph 0->1(4),0->2(1),2->1(2),1->3(1),2->3(5), print the actual shortest path from 0 to 3 `0 2 1 3`. A parallel prev slice filled during relaxation lets us backtrack from the target.

## Run

    go run .
