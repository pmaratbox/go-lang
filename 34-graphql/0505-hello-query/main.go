package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	// Schema: type Query { hello: String }, resolver returns "world".
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: fields}),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Execute "{ hello }" in-process (no HTTP server).
	r := graphql.Do(graphql.Params{Schema: schema, RequestString: "{ hello }"})
	if len(r.Errors) > 0 {
		log.Fatal(r.Errors)
	}

	// Extract the resolved value from the execution result's data.
	data := r.Data.(map[string]interface{})
	fmt.Println(data["hello"])
}
