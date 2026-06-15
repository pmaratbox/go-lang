# 0547 — Traffic light

Uses the `github.com/looplab/fsm` finite-state-machine library. Three states `red`, `green`, and `yellow` are cycled by a single `next` event defined once per source state (`red->green`, `green->yellow`, `yellow->red`). Starting in `red`, firing `next` twice walks `red -> green -> yellow`, so `f.Current()` reports `yellow`, which is printed lowercased. The final state comes from the machine, not a hardcoded value.

## Run

    go run .
