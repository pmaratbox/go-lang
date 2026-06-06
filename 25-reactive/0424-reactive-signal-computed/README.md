# 0424 — Signal + Computed

Implement fine-grained reactivity: a writable signal and a derived computed that recomputes when its dependency changes. In Go, the recompute callback is a closure captured as a subscriber, keeping the dependency wiring explicit and synchronous.

## Run

    go run .
