package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

func main() {
	fields := graphql.Fields{
		"greet": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				return "hello " + name, nil
			},
		},
	}
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: fields}),
	})
	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `{ greet(name: "alice") }`,
	})
	fmt.Println(r.Data.(map[string]interface{})["greet"])
}
