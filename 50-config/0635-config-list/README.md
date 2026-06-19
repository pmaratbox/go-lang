# 0635 — List value

Uses the real configuration library [spf13/viper](https://github.com/spf13/viper).
It loads the fixed `config.json` via `viper.SetConfigFile` + `viper.ReadInConfig`
and reads the array key `hosts` with `viper.GetStringSlice`, joining the elements
with commas to print the extracted value `a,b,c`.

## Run

    go run .
