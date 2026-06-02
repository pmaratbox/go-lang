# 0027 — File I/O

Write `hello, file` to a file, read it back, delete the file, and print `read: hello, file`. `os.WriteFile` and `os.ReadFile` handle opening, transferring, and closing in one call, each returning an `error` that is checked explicitly; `os.Remove` deletes the file. `ReadFile` returns a `[]byte`, which `%s` prints as text.

## Run

    go run .
