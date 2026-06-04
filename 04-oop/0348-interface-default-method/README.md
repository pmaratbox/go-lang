# 0348 — Interface Default Method

Define an interface with a default greet() returning "hi" and an implementer that overrides it to "hey", printing `hi hey`. Go has no interface defaults, so an embedded struct supplies the default and a redefined method overrides it.

## Run

    go run .
