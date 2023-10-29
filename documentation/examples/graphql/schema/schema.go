package schema

import (
	"getting-started/objects"

	"github.com/graphql-go/graphql"
)

// ? Function that returns a new schema
func CreateSchema() (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query: objects.Query,
		},
	)
}
