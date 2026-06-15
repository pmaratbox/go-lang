# 0511 — Mutation

Define a GraphQL schema with `github.com/graphql-go/graphql` that includes a `Mutation` type whose `addUser(name: String!): User` field constructs and returns a `User`. The mutation `mutation { addUser(name: "bob") { name } }` is executed in-process with `graphql.Do` (no HTTP server). The program extracts `data.addUser.name` from the execution result and prints it — `bob`.

## Run

    go run .
