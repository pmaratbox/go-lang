# 0397 — Log Level Filter

With a threshold of WARN, log messages at INFO, WARN, and ERROR but only emit WARN and ERROR, on two lines. An ordered iota enum makes the threshold comparison a simple integer check.

## Run

    go run .
