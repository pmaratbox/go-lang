# 0652 — Return a value

Uses the real retry library
[cenkalti/backoff/v4](https://github.com/cenkalti/backoff). A scripted
operation fails once and then returns the string `ok`. `backoff.RetryWithData`
wraps it with `WithMaxRetries(&backoff.ZeroBackOff{}, 4)`, allowing up to four
zero-delay retries, and the library carries back the value produced on the
successful attempt. The program prints that returned value (not the attempt
count), so the output is `ok`.

## Run

    go run .
