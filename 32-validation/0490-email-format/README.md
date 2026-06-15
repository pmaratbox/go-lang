# 0490 — Email format

Uses the `github.com/go-playground/validator/v10` library to enforce the
`email` constraint on the `Email` field. The struct tag `validate:"email"` makes
`v.Struct` reject any value that is not a syntactically valid email address;
here the input `{Email: "not-an-email"}` fails. The program iterates the
returned `validator.ValidationErrors`, lowercases each `Field()`, sorts them,
and prints the failing field name(s) — output is `email`, never the library's
message text.

## Run

    go run .
