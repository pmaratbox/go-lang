# 0633 — Read a value

Uses the real configuration library [spf13/viper](https://github.com/spf13/viper).
It loads the fixed `config.json` via `viper.SetConfigFile` + `viper.ReadInConfig`
and reads the top-level string key `name` with `viper.GetString`, printing the
extracted value `myapp`.

## Run

    go run .
