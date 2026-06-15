# 0473 — Multiple values

Uses the `github.com/spf13/cobra` CLI library to collect a repeated option into a
list via `IntSliceVar`, which binds the `--num` flag so it may appear multiple
times. The parser runs against a fixed, hardcoded argv
(`--num 1 --num 2 --num 3`) supplied through `cmd.SetArgs`, so the output is
deterministic regardless of the real process arguments. The collected values are
summed and printed.

## Run

    go run .
