# 0522 — Add days

Using Go's standard `time` package, this parses the fixed date `2026-06-15`
with `time.Parse`, then adds 10 days using `Time.AddDate(0, 0, 10)`, the
library's calendar arithmetic that correctly rolls over month boundaries. The
result is formatted back to ISO `2006-01-02` layout.

## Run

    go run .
