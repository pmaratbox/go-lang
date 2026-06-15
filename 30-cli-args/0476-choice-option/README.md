# 0476 — Choice option

Uses the `github.com/spf13/cobra` CLI library with a custom `pflag.Value`
to implement a choice/enum option: the `--color` flag is restricted to the
set `{red, green, blue}`, and `Set` rejects anything else. The program parses
a fixed argv `["--color", "green"]` (not the real process arguments) so the
output is deterministic, then prints the chosen value.

## Run

    go run .
