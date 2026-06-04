# 0125 — Temp File Roundtrip

Write a string to a temporary file, read it back, confirm it matches, delete the file, and print `roundtrip: ok`. The `os.CreateTemp` helper opens a uniquely named file so concurrent runs never collide.

## Run

    go run .
