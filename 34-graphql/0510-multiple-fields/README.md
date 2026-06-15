# 0510 — Multiple fields

Build a GraphQL schema with the `github.com/graphql-go/graphql` library — `type User { name: String age: Int }` and `type Query { user: User }` whose resolver returns one user — then execute the query `{ user { name age } }` entirely in-process with `graphql.Do` (no HTTP server). A selection set can request multiple fields of an object at once; both resolved values are read out of the execution result's `Data` map and printed, `name` then `age`, each on its own line.

## Run

    go run .
