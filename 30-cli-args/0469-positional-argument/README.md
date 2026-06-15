# 0469 — Positional argument

Define a CLI with one positional argument `name` using Go's real CLI library `github.com/spf13/cobra`. The command declares `Args: cobra.ExactArgs(1)` to require exactly one positional, and reads it from `args[0]` inside `Run`. For determinism the program parses a fixed argv `["alice"]` via `cmd.SetArgs(...)` instead of the real process arguments, so running with no arguments always prints `alice`.

## Run

    go run .
