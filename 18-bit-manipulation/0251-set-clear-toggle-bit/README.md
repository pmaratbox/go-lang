# 0251 — Set, Clear, Toggle Bit

On bit position 1: set it on 0 (->2), clear it on 2 (->0), toggle it on 0 (->2), printing `2 0 2`. Go's `&^` operator clears bits directly without a separate NOT.

## Run

    go run .
