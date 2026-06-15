# 0474 — Subcommand

Define a subcommand `add` with the `github.com/spf13/cobra` library and dispatch to it. The `add` subcommand takes two integer positionals (declared via `cobra.ExactArgs(2)`); when invoked it sums them and prints the result. To stay deterministic the parser is fed a fixed argv (`root.SetArgs([]string{"add", "2", "3"})`) instead of the real process arguments, so running with no arguments always prints `5`.

## Run

    go run .
