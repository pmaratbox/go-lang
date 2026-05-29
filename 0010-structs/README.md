# 0010 — Structs

Define a `Person` struct with a `Name` and an `Age`, create one ("Ada", 36),
and print each field. `type Person struct { ... }` declares the type; a
composite literal `Person{Name: ..., Age: ...}` builds a value. Fields starting
with an uppercase letter are exported (visible from other packages).

## Run

    go run .
