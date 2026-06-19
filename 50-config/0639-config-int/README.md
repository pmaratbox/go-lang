# 0639 — Integer value

Uses the real configuration library [spf13/viper](https://github.com/spf13/viper).
It loads the fixed `config.json` via `viper.SetConfigFile` + `viper.ReadInConfig`
and reads the integer key `retries` with `viper.GetInt`, printing the extracted
value `3`.

## Run

    go run .
