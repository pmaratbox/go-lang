# 0330 — Lens Get/Set

Use a lens over the nested value {a:{b:1}} to get b (1) and to set b to 2, printing `1 2`. Go's value semantics make the setter return a copy, leaving the original struct untouched.

## Run

    go run .
