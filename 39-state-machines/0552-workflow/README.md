# 0552 — Workflow

Uses the `github.com/looplab/fsm` finite-state-machine library to model a multi-step approval workflow. States and transitions are declared as `fsm.Events` ({Name, Src, Dst}); the machine starts in `idle`, and firing `submit` then `approve` drives it `idle -> pending -> approved`. `f.Current()` reports the resulting state, lowercased for output.

## Run

    go run .
