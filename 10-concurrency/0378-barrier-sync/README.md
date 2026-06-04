# 0378 — Barrier Synchronization

Have 3 threads each arrive at a barrier before any proceeds, then print `all reached: 3`. A WaitGroup used as a barrier lets each goroutine signal arrival and wait for the others.

## Run

    go run .
