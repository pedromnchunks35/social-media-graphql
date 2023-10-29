package objects

import (
	"getting-started/database"

	"github.com/graphql-go/graphql"
)

// ? Create a object of type query
var Query = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: UserType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, isOK := p.Args["id"].(int)
					if isOK {
						return (*database.Data)[id], nil
					}
					return nil, nil
				},
			},
		},
	},
)
