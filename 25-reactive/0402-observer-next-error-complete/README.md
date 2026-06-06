# 0402 — Observer Contract

Demonstrate the observer contract next*-then-terminal: emit 1 and 2, complete, and show that a post-complete next is ignored. A `stopped` bool guards the methods so post-terminal calls return early as idiomatic no-ops.

## Run

    go run .
