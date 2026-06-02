# 0060 — Priority Queue

Push `3`, `1`, and `2` into a min-priority-queue, then pop them all and print them in priority (ascending) order: `1 2 3`. Go's `container/heap` supplies the algorithm; you provide a type implementing `heap.Interface` (`Len`/`Less`/`Swap`/`Push`/`Pop`). `Less` returning `<` makes it a min-heap.

## Run

    go run .
