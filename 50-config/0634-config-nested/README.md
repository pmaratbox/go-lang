# 0634 — Nested key

Uses the real configuration library [spf13/viper](https://github.com/spf13/viper).
It loads the fixed `config.json` via `viper.SetConfigFile` + `viper.ReadInConfig`
and reads the nested integer key `server.port` with `viper.GetInt` (dot-path
access into the nested object), printing the extracted value `8080`.

## Run

    go run .
