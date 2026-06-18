# 0574 — TOML array of tables

Parse a TOML array-of-tables (`[[servers]]`) with the `github.com/pelletier/go-toml/v2`
library. `toml.Unmarshal` decodes the document into a `map[string]any`, where the
`servers` key holds a `[]any` of `map[string]any` entries (one per `[[servers]]`
block). Each server's `name` is collected and joined with commas to print
`alpha,beta`.

## Run

    go run .
