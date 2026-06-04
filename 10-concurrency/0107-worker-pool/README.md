# 0107 — Worker Pool

Distribute squaring of 1..4 across a pool of workers, collect the results, and print them sorted ascending `1 4 9 16`. A pool of goroutines reads jobs from one channel and writes squares to a results channel.

## Run

    go run .
