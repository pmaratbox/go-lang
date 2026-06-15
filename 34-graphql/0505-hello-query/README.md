# 0505 — Hello query

Build a GraphQL schema with the `github.com/graphql-go/graphql` library — a single `type Query { hello: String }` whose resolver returns `"world"` — then execute the query `{ hello }` entirely in-process with `graphql.Do` (no HTTP server). The resolved value is read out of the execution result's `Data` map (`r.Data.(map[string]interface{})["hello"]`) and printed: `world`.

## Run

    go run .
