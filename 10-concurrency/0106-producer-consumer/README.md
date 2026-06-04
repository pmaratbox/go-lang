# 0106 — Producer / Consumer

A producer sends 1..5 to a consumer that sums them, printing `15`. A buffered channel decouples the producer goroutine from the consuming main loop.

## Run

    go run .
