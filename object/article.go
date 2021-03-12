package object

import "github.com/graphql-go/graphql"

// type Article struct {
// 	ID        int64     `json:"id"`
// 	Title     string    `json:"title"`
// 	Subtitle  string    `json:"subtitle"`
// 	Content   string    `json:"content"`
// 	CreatedAt time.Time `json:"createdAt"`
// 	UserID    int64     `json:"userId"`
// }

// NewArticleType is a constructor of graphql article type.
func NewArticleType() (userType *graphql.Object) {
	userType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Article",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"subtitle": &graphql.Field{
					Type: graphql.String,
				},
				"content": &graphql.Field{
					Type: graphql.String,
				},
				"createdAt": &graphql.Field{
					Type: graphql.DateTime,
				},
				"userId": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)

	return
}

// NewArticleFieldConfigArgument is a constructor of article field config argument.
func NewArticleFieldConfigArgument() (userFieldConfigArgument graphql.FieldConfigArgument) {
	userFieldConfigArgument = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"userId": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}

	return
}
