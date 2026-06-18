# 0572 — TOML table

Parse a TOML document with the third-party `github.com/pelletier/go-toml/v2`
library. `toml.Unmarshal` decodes the `[server]` table into a
`map[string]any`, where the table value is itself a `map[string]any`. We read
`server.host` (a string) and `server.port` (an integer) and print them in a
fixed `host=<host> port=<port>` format.

## Run

    go run .
