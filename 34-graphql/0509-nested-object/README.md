# 0509 — Nested object

Build a GraphQL schema with the `github.com/graphql-go/graphql` library where a `Query.user` field returns a `User` object that itself nests an `Address` object. The query `{ user { address { city } } }` is executed in-process with `graphql.Do` (no HTTP server), and the resolved value is pulled out of `result.Data` by walking the nested maps — `data.user.address.city` — which prints `oslo`.

## Run

    go run .
