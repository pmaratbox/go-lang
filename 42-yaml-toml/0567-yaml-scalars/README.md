# 0567 — YAML scalars

Parse a YAML mapping with the `gopkg.in/yaml.v3` library and read its scalar
fields. `yaml.Unmarshal` decodes the fixed document
`name: Alice\nrole: admin\nage: 30\n` into a `map[string]any`, where `name` and
`role` become strings and `age` becomes an integer. The three values are printed
space-joined as `Alice admin 30`.

## Run

    go run .
