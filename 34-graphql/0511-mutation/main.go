package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	// type User { name: String }
	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"name": &graphql.Field{Type: graphql.String},
		},
	})

	// type Query { _: String } — a schema needs a Query type even when
	// the lesson only exercises the mutation.
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"_": &graphql.Field{Type: graphql.String},
		},
	})

	// type Mutation { addUser(name: String!): User }
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, _ := p.Args["name"].(string)
					return map[string]interface{}{"name": name}, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Execute the mutation in-process (no HTTP server).
	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `mutation { addUser(name: "bob") { name } }`,
	})
	if len(r.Errors) > 0 {
		log.Fatal(r.Errors)
	}

	// Extract the resolved value from the execution result's data.
	data := r.Data.(map[string]interface{})
	addUser := data["addUser"].(map[string]interface{})
	fmt.Println(addUser["name"])
}
