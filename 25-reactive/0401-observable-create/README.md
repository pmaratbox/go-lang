# 0401 — Create an Observable

Build a push-based Observable from scratch that emits 1, 2, 3 to its observer and then completes. In Go an Observable is modeled as a function type that pushes values to an Observer struct of `next`/`complete` closures.

## Run

    go run .
