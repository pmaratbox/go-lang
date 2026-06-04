# 0151 — Singleton

Obtain a singleton instance twice and confirm both references are the same object, printing `same: yes`. `sync.Once` lazily initializes the shared instance exactly once, even across goroutines.

## Run

    go run .
