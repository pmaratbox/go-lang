# 0571 — TOML scalars

Parse a small TOML document with the `github.com/pelletier/go-toml/v2`
library. `toml.Unmarshal` decodes `title = "demo"\nversion = 2\n` into a
`map[string]any`, where `title` becomes a string and `version` becomes an
integer. The two top-level scalar values are printed space-joined as
`demo 2` (the version prints plainly as `2`).

## Run

    go run .
