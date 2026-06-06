# 0411 — Concat Streams

Implement concat: subscribe to the second source only after the first completes; concat [1,2] then [3,4]. In Go, observers are small structs of next/complete closures, so concat just wires the first source's complete to a subscription on the second.

## Run

    go run .
