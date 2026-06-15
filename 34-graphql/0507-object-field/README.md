# 0507 — Object field

Use the `github.com/graphql-go/graphql` library to define a schema with a custom object type `User { name: String }` exposed through `Query { user: User }`. The `user` resolver returns a map standing in for a `User` value, and the query `{ user { name } }` selects a single field from that nested object. The query is executed entirely in-process with `graphql.Do` (no HTTP server); the program extracts `data.user.name` from the execution result and prints it — `alice`.

## Run

    go run .
