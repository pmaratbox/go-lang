# 0380 — Compare-And-Swap Loop

Increment a shared value to 100 using a CAS retry loop from multiple threads, printing `100`. `atomic.CompareAndSwapInt64` in a retry loop guards each lock-free increment.

## Run

    go run .
