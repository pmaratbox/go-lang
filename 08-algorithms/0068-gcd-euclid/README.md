# 0068 — GCD (Euclid)

Compute the greatest common divisor of `48` and `36` with Euclid's algorithm (repeatedly replace the pair with `(b, a % b)` until the remainder is zero) and print it: `12`. Go's multiple assignment `a, b = b, a%b` evaluates the right side first, so both update simultaneously until `b` is zero.

## Run

    go run .
