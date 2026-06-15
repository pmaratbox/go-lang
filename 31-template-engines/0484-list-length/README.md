# 0484 — List length

Render the length of a list using Go's real `text/template` engine. The template `{{len .Items}}` calls the built-in `len` function on the `Items` slice `[1, 2, 3]` supplied as data, so the engine evaluates and prints the slice length `3`.

## Run

    go run .
