# 0636 — Boolean value

Uses the real configuration library [spf13/viper](https://github.com/spf13/viper).
It loads the fixed `config.json` via `viper.SetConfigFile` + `viper.ReadInConfig`
and reads the boolean key `debug` with `viper.GetBool`, printing it lowercase as
`true`.

## Run

    go run .
