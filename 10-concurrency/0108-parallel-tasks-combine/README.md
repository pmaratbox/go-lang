# 0108 — Parallel Tasks Combined

Run two independent tasks that produce 10 and 20 concurrently, then combine (sum) their results into `30`. Each task runs in its own goroutine and delivers its result over a channel.

## Run

    go run .
