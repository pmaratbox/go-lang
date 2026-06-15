package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

func main() {
	itemType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Item",
		Fields: graphql.Fields{
			"id": &graphql.Field{Type: graphql.Int},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"item": &graphql.Field{
				Type: itemType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					return map[string]interface{}{"id": id}, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{Query: queryType})

	query := `query($id: Int!) { item(id: $id) { id } }`

	r := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query,
		VariableValues: map[string]interface{}{"id": 42},
	})

	item := r.Data.(map[string]interface{})["item"].(map[string]interface{})
	fmt.Println(item["id"])
}
