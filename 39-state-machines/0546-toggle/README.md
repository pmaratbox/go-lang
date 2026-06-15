# 0546 — Toggle

Uses the `github.com/looplab/fsm` finite-state-machine library. Two states `off` and `on` are wired with a single `toggle` event whose source/destination pair flips the machine between them (`off->on` and `on->off`). Starting in `off`, firing `toggle` three times walks `off -> on -> off -> on`, so `f.Current()` reports `on`, which is printed lowercased.

## Run

    go run .
