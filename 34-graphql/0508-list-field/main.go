package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

func main() {
	fields := graphql.Fields{
		"numbers": &graphql.Field{
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return []int{1, 2, 3}, nil
			},
		},
	}
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: fields}),
	})
	r := graphql.Do(graphql.Params{Schema: schema, RequestString: "{ numbers }"})
	data := r.Data.(map[string]interface{})
	for _, n := range data["numbers"].([]interface{}) {
		fmt.Println(n)
	}
}
