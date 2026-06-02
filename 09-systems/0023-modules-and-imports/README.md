# 0023 — Modules & Imports

Define `square(n)` in a separate `mathutil` module and import it from the main program, printing `square(8) = 64` across the module boundary. Code is organized into packages: `mathutil/mathutil.go` declares `package mathutil`, and `main` imports it by its full module path. Only identifiers starting with an uppercase letter (`Square`) are exported across the package boundary, and `go run .` builds the whole module.

## Run

    go run .
