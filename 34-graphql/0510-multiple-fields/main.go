package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	// type User { name: String age: Int }
	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"name": &graphql.Field{Type: graphql.String},
			"age":  &graphql.Field{Type: graphql.Int},
		},
	})

	// type Query { user: User }, resolver returns one user.
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{"name": "alice", "age": 30}, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: queryType})
	if err != nil {
		log.Fatal(err)
	}

	// Execute "{ user { name age } }" in-process (no HTTP server).
	r := graphql.Do(graphql.Params{Schema: schema, RequestString: "{ user { name age } }"})
	if len(r.Errors) > 0 {
		log.Fatal(r.Errors)
	}

	// Extract multiple selected fields of the user object from the result data.
	user := r.Data.(map[string]interface{})["user"].(map[string]interface{})
	fmt.Println(user["name"])
	fmt.Println(user["age"])
}
