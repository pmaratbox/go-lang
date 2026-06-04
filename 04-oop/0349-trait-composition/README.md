# 0349 — Trait Composition

Compose two capabilities (A printing "a", B printing "b") into one type and invoke both, printing `a b`. Go composes behavior by embedding both structs, promoting their methods onto the host type.

## Run

    go run .
