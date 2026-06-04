# 0109 — Run-Once Initialization

Ensure an initializer runs exactly once even when several threads race to trigger it, printing `init count: 1`. sync.Once guarantees the init body runs a single time regardless of how many goroutines call Do.

## Run

    go run .
