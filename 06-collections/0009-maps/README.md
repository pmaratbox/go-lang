# 0009 — Maps

Build a map `map[string]int{...}`, look up `"two"`, and print its value and the
map's size. `m[key]` returns the value, or the value type's zero value if the
key is absent — the comma-ok form `v, ok := m[key]` distinguishes "present but
zero" from "missing". `len(m)` counts entries. Map iteration order is
randomized.

## Run

    go run .
