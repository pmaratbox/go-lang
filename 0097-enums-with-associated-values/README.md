# 0097 — Enums with Associated Values

Define a shape type carrying associated data — `Rect(2, 3)` and `Square(4)` — compute each area by matching on the variant, and print `6` and `16`. Go has no sum types with payloads; an interface implemented by each variant struct provides the polymorphic `area`.

## Run

    go run .
