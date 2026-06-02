# 0057 — Sieve of Eratosthenes

Use the Sieve of Eratosthenes to find every prime number up to `10` and print them: `2 3 5 7`. A `[]bool` (zero value `false`) is set true for `2..n`, then multiples are struck from `i*i`; the guard `i*i <= n` bounds the outer loop without a square root.

## Run

    go run .
