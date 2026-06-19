# 0637 — Default for missing key

Uses the real configuration library [spf13/viper](https://github.com/spf13/viper).
It loads the fixed `config.json` via `viper.SetConfigFile` + `viper.ReadInConfig`,
then registers a default for a key absent from the file using
`viper.SetDefault("missing", "fallback")`. Reading `missing` with
`viper.GetString` resolves to the default, printing `fallback`.

## Run

    go run .
