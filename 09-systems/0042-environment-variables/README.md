# 0042 — Environment Variables

Read the environment variable `LESSON_ENV_VAR`, falling back to `default` when it is unset, and print `value: default`. `os.LookupEnv` returns the value and a boolean reporting whether it was set, so an unset variable is distinguished from an empty one; `os.Getenv` just returns `""` for both.

## Run

    go run .
