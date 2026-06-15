# 0551 — Reset

Uses the `github.com/looplab/fsm` finite-state-machine library. We declare two states (`idle`, `running`) and two transitions: `start` moves `idle -> running`, and `reset` moves `running -> idle`. Starting in `idle`, firing `start` then `reset` walks the machine back to its initial state, demonstrating how a reset event returns the FSM to where it began. The resulting state is read from `f.Current()` and printed lowercased.

## Run

    go run .
