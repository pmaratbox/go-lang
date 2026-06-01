# 0031 — Type Conversion & Parsing

Parse the string `"42"` into an integer and `"3.5"` into a float, then convert the integer back to a string, printing `int: 42`, `float: 3.5`, and `str: 42`. The `strconv` package converts explicitly: `Atoi` for ints, `ParseFloat(s, 64)` for floats, and `Itoa` back to a string. Each parser returns a value *and* an `error`, so failures are values to be checked, not exceptions.

## Run

    go run .
