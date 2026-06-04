# 0382 — Fork-Join Sum

Recursively fork the sum of [1..8] into halves and join the partial sums, printing `36`. Each recursive split runs the left half in a goroutine and joins it via a buffered channel.

## Run

    go run .
