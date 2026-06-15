# 0548 — Guarded transition

Uses the `github.com/looplab/fsm` finite-state-machine library to model a door whose `open` event is *guarded*: it is only valid from the `unlocked` state. The guard is expressed through the transition's `Src` list — `fsm.NewFSM` rejects any event fired from a state not listed in `Src`, leaving the machine unchanged. Starting from `locked`, firing `unlock` then `open` lands the machine in `open`, which we read back from `f.Current()` and lowercase.

## Run

    go run .
