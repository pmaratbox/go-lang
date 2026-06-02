# 0039 — Records & Value Equality

Create two points with the same fields, print one as `point: (1, 2)`, and compare them by value to print `equal: yes`. Go structs are comparable out of the box: `==` does a field-by-field value comparison when every field is comparable. No annotation or generated method is needed.

## Run

    go run .
