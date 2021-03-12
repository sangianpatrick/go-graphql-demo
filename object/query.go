package object

import (
	"github.com/graphql-go/graphql"
)

// NewQueryType is a constructor of query type.
func NewQueryType() (queryType *graphql.Object) {
	queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: graphql.Fields{},
		},
	)

	return
}
