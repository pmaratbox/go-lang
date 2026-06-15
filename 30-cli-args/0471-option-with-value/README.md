# 0471 — Option with value

Uses the `github.com/spf13/cobra` CLI library to define an option that takes a
value: a `--name` string flag bound with `Flags().StringVar`. Instead of
reading the real process arguments, the command parses a fixed argv
(`["--name", "alice"]`) via `cmd.SetArgs`, so running the program with no
arguments always produces the same deterministic output. The parsed value is
printed.

## Run

    go run .
