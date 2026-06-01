# 0054 — Resource Cleanup & Defer

Acquire a resource, use it, and let the language release it automatically at scope exit, printing `open`, `use`, and `close` in that order. `defer` schedules a call to run when the surrounding function returns; multiple defers run last-in-first-out. It is Go's standard cleanup idiom (e.g. `defer file.Close()`).

## Run

    go run .
