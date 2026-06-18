# 0573 — TOML array

Parse a TOML array with the `github.com/pelletier/go-toml/v2` library.
`toml.Unmarshal` decodes `tags = ["red", "green", "blue"]\n` into a
`map[string]any`, where the `tags` key holds a `[]any` of strings. The
elements are joined with commas via `strings.Join` to print
`red,green,blue`.

## Run

    go run .
