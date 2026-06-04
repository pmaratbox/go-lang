# 0322 — Lazy Filter + Take

From a lazy stream of naturals, filter the even ones and take three, printing `2 4 6`. Each stage is a goroutine wired by channels, so the pipeline pulls only what `take` consumes.

## Run

    go run .
