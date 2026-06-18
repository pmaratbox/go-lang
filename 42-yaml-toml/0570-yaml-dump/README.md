# 0570 — Dump YAML

Serialize a map to sorted block-style YAML with the `gopkg.in/yaml.v3`
library. A `map[string]any` holding `name=Alice` (string), `age=30`
(integer), and `city=Paris` (string) is passed to `yaml.Marshal`, which
emits the keys in sorted order using block style — no flow braces and no
quotes on these simple scalars. The integer prints plainly as `30`, and
the bytes are written straight to stdout with `os.Stdout.Write`.

## Run

    go run .
