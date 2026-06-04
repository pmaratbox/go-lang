# 0105 — Channels / Message Passing

Send the values 1, 2, 3 through a channel (or queue) from one thread and receive them in order, printing `1 2 3`. A goroutine sends on a channel and closes it; the main loop ranges to receive in order.

## Run

    go run .
