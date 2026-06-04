# 0208 — Virtual Proxy

Use a lazy virtual proxy that loads the real subject only on first access, printing `loaded`. A nil pointer guard inside the method lazily constructs the real subject on demand.

## Run

    go run .
