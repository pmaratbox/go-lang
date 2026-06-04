# 0103 — Mutex-Protected Counter

Have multiple threads each increment a shared counter under a mutex so the total is exactly `1000`. A sync.Mutex guards every increment against data races.

## Run

    go run .
