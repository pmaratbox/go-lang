# 0415 — SwitchMap

Implement switchMap: when a new outer value arrives, cancel the previous inner subscription before starting the new one. In Go we keep the previous inner's scheduled tasks in a slice and mark them dead via the virtual-time scheduler before scheduling the next inner.

## Run

    go run .
