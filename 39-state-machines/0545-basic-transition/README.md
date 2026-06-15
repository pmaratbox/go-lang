# 0545 — Basic transition

Uses the `github.com/looplab/fsm` finite-state-machine library. `fsm.NewFSM` defines an initial state plus a set of `Events`, where each event names a transition from a source state (`Src`) to a destination state (`Dst`). The turnstile starts `locked`; firing the `coin` event drives the deterministic transition to `unlocked`, and `f.Current()` reports the machine's resulting state.

## Run

    go run .
