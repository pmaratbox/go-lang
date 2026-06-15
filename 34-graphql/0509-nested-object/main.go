package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

func main() {
	addressType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Address",
		Fields: graphql.Fields{
			"city": &graphql.Field{Type: graphql.String},
		},
	})

	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"address": &graphql.Field{Type: addressType},
		},
	})

	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{
						"address": map[string]interface{}{
							"city": "oslo",
						},
					}, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{Query: query})

	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: "{ user { address { city } } }",
	})

	data := r.Data.(map[string]interface{})
	user := data["user"].(map[string]interface{})
	address := user["address"].(map[string]interface{})
	fmt.Println(address["city"])
}
