# 0417 — BehaviorSubject

Implement a BehaviorSubject that holds a current value and replays it immediately to each new subscriber. Observers are plain `func(int)` closures appended to a slice and invoked synchronously on subscribe and on each `Next`.

## Run

    go run .
