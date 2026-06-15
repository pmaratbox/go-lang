# 0506 — Field argument

Pass an argument to a field with Go's real `github.com/graphql-go/graphql` library. The schema declares `Query.greet(name: String!): String`, where the `Args` map gives the field a non-null `name` argument; the resolver reads it from `p.Args["name"]` and returns `"hello " + name`. The query `{ greet(name: "alice") }` is executed in-process with `graphql.Do` (no HTTP server), and the resolved value is extracted from `r.Data` to print `hello alice`.

## Run

    go run .
