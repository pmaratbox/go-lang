# 0215 — Chain of Responsibility

Pass a request of level 2 along a handler chain so the level-2 handler handles it, printing `handled by 2`. Each handler holds a pointer to the next and forwards requests it cannot handle.

## Run

    go run .
