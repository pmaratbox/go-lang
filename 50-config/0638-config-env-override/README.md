# 0638 — Env override

Uses the real configuration library [spf13/viper](https://github.com/spf13/viper).
It loads the fixed `config.json`, then binds the `name` key to the in-process
`NAME` environment variable via `viper.BindEnv`. Viper resolves bound env vars
with higher priority than the file, so `viper.GetString("name")` returns the
overriding value `from-env` instead of the file's `myapp`.

## Run

    go run .
