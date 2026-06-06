# 0404 — Cold vs Hot Observable

Contrast a cold observable (re-runs its producer per subscriber) with a hot one (shares a single execution, so late subscribers miss earlier values). In Go we model each observer as a struct holding a `next` closure, and the hot observable shares one slice of observers via a pointer.

## Run

    go run .
