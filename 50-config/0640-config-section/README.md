# 0640 — Read a section

Uses the real configuration library [spf13/viper](https://github.com/spf13/viper).
It loads the fixed `config.json` via `viper.SetConfigFile` + `viper.ReadInConfig`
and reads two keys from the `server` section with `viper.GetString("server.host")`
and `viper.GetInt("server.port")`, printing them as `host:port` -> `0.0.0.0:8080`.

## Run

    go run .
