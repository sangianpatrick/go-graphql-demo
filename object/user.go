package object

import (
	"github.com/graphql-go/graphql"
)

// NewUserType is a constructor of graphql user type.
func NewUserType() (userType *graphql.Object) {
	userType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"createdAt": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	)

	return
}

// NewUserFieldConfigArgument is a constructor of user field config argument.
func NewUserFieldConfigArgument() (userFieldConfigArgument graphql.FieldConfigArgument) {
	userFieldConfigArgument = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}

	return
}
