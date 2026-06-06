# 0418 — ReplaySubject

Implement a ReplaySubject with a buffer of the last 2 values, replayed to a late subscriber, which then also receives new values. A slice trimmed with reslicing keeps the most recent values that are pushed to each new observer.

## Run

    go run .
