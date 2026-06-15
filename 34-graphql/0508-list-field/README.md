# 0508 — List field

Build a GraphQL schema with the `github.com/graphql-go/graphql` library where `type Query` has a list field `numbers: [Int]` (a list of scalars, declared via `graphql.NewList(graphql.Int)`) whose resolver returns `[]int{1, 2, 3}`. The query `{ numbers }` is executed entirely in-process with `graphql.Do` (no HTTP server). The resolved list is read out of the execution result's `Data` map (`r.Data.(map[string]interface{})["numbers"].([]interface{})`) and each element is printed on its own line.

## Run

    go run .
