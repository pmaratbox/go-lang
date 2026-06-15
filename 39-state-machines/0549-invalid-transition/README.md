# 0549 — Invalid transition

Using the [`github.com/looplab/fsm`](https://github.com/looplab/fsm) library, a turnstile starts in `locked`. Firing `push` — which has no transition defined from `locked` — is rejected: `Event` returns an `InvalidEventError` and the machine leaves its current state unchanged. We ignore the error and print the state, still `locked`.

## Run

    go run .
