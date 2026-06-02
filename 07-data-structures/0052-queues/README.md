# 0052 — Queues

Enqueue `1`, `2`, and `3` into a queue, then dequeue them all and print them in first-in-first-out order: `1 2 3`. Go has no queue type; a slice is the idiom — `append` to enqueue and reslice `queue[1:]` to drop the front. `container/list` offers a linked list when front removal must stay O(1).

## Run

    go run .
