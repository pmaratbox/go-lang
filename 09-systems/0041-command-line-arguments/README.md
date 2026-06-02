# 0041 — Command-line Arguments

Read the first command-line argument and greet it, so running with `Ada` prints `hello, Ada`. `os.Args` holds the arguments with `os.Args[0]` as the program path, so the first user argument is `os.Args[1]`. The `flag` package parses named options.

## Run

    go run . Ada
