# 0475 — Required option

Declare a CLI option as required using the `github.com/spf13/cobra` library. The integer flag `--id` is registered with `cmd.Flags().IntVar` and then marked mandatory via `cmd.MarkFlagRequired("id")`, so cobra errors out if it is missing. For deterministic output the program parses a fixed argv `["--id", "42"]` passed through `cmd.SetArgs` rather than the real process arguments, printing `42`.

## Run

    go run .
