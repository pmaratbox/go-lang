# 0569 — YAML sequence

Parse a YAML document with a block sequence using Go's real YAML library
`gopkg.in/yaml.v3`. `yaml.Unmarshal` decodes `fruits:\n  - apple\n  - banana\n
  - cherry\n` into a `map[string]any`, where the `fruits` key holds an `[]any`
of strings. We type-assert each element and join them with commas to produce
`apple,banana,cherry`.

## Run

    go run .
