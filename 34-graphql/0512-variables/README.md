# 0512 — Query variables

Build a GraphQL schema with the `github.com/graphql-go/graphql` library — `type Query { item(id: Int!): Item }` where `type Item { id: Int }` — then execute the query `query($id: Int!) { item(id: $id) { id } }` in-process with `graphql.Do` (no HTTP server). The `$id` variable is supplied through the execution's `VariableValues` map (`{"id": 42}`), not interpolated into the query string; the resolver echoes it back, and `data.item.id` is read out of the result and printed: `42`.

## Run

    go run .
