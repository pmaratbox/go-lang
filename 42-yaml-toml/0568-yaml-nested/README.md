# 0568 — Nested YAML mapping

Parse a nested YAML mapping with the `gopkg.in/yaml.v3` library. `yaml.Unmarshal`
decodes `server:\n  host: localhost\n  port: 8080\n` into a `map[string]any`,
where the `server` key holds a nested `map[string]any`. We read `server.host`
and `server.port` and print them joined as `host:port` -> `localhost:8080`.

## Run

    go run .
