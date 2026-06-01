# 0037 — Operator Overloading

Define how `+` (or an `add` method) combines two points, then add `(1, 2)` and `(3, 4)` and print `(4, 6)`. Go deliberately omits operator overloading; a named `Add` method does the work. Implementing `String() string` (the `Stringer` interface) lets `fmt.Println` print `(4, 6)`.

## Run

    go run .
