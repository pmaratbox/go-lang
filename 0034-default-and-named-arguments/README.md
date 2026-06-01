# 0034 — Default & Named Arguments

Give a `greet` function a default greeting, then call it once without the greeting and once overriding it, printing `Hello, Ada` and `Hi, Ada`. Go has neither default nor named arguments. The idioms are a wrapper function that supplies the default (shown here), a variadic `...Option` functional-options pattern, or passing a config struct.

## Run

    go run .
