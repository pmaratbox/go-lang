# 0550 — Transition count

Uses the `github.com/looplab/fsm` finite-state-machine library. The machine cycles through `green -> yellow -> red -> green`, and an `"after_event"` callback fires after every successful transition, incrementing a counter. Firing three valid `next` events therefore drives the counter to `3` — the value comes from the per-transition action, not a hardcoded literal.

## Run

    go run .
