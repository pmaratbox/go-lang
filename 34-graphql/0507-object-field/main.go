package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

func main() {
	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"name": &graphql.Field{Type: graphql.String},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{"name": "alice"}, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{Query: queryType})

	r := graphql.Do(graphql.Params{Schema: schema, RequestString: "{ user { name } }"})

	user := r.Data.(map[string]interface{})["user"].(map[string]interface{})
	fmt.Println(user["name"])
}
