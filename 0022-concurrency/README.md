# 0022 — Concurrency

Start two tasks that produce `1` and `2`, let them run concurrently, then join their results and print `sum: 3`. `go task(...)` launches a goroutine — a lightweight thread the runtime multiplexes onto OS threads. The two goroutines send into a buffered `chan int`, and the main goroutine receives twice (`<-ch + <-ch`), which also synchronizes. Channels are Go's primary way for goroutines to communicate.

## Run

    go run .
