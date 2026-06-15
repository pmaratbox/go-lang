# 0472 — Default value

Uses the [spf13/cobra](https://github.com/spf13/cobra) CLI library to define an
integer flag `--count` with a default value of `1`. The command parses a fixed,
hardcoded empty argv (`[]`) instead of the real process arguments, so the run is
deterministic: since `--count` is absent, cobra leaves it at its default and the
program prints `1`.

## Run

    go run .
