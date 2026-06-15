# 0470 — Boolean flag

Define a boolean flag `--verbose` on a `cobra.Command` from the `github.com/spf13/cobra` CLI library and read its parsed value. To stay deterministic the command parses a hardcoded argv (`[]string{"--verbose"}`) set via `cmd.SetArgs` instead of the real process arguments, so running with no arguments still prints the flag value `true` (lowercase).

## Run

    go run .
