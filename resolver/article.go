package resolver

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/sangianpatrick/go-graphql-demo/entity"
)

// ArticleResolver is a collection of behavior of article resolver.
type ArticleResolver interface {
	GetArticleByID(rp graphql.ResolveParams) (data interface{}, err error)
	GetManyArticleByUserID(rp graphql.ResolveParams) (data interface{}, err error)
}

type articleResolverImplementation struct {
}

// NewArticleResolver is a constructor.
func NewArticleResolver() ArticleResolver {
	return &articleResolverImplementation{}
}

func (r articleResolverImplementation) GetArticleByID(rp graphql.ResolveParams) (data interface{}, err error) {
	id := rp.Args["id"].(int)
	if id != 1 {
		return
	}
	var article = entity.Article{
		ID:        1,
		Title:     "Golang Graphql Demo (Clean Code)",
		Subtitle:  "Build simple graphql application",
		Content:   "No Content",
		CreatedAt: time.Now().In(time.UTC),
	}

	data = article

	return
}

func (r articleResolverImplementation) GetManyArticleByUserID(rp graphql.ResolveParams) (data interface{}, err error) {
	userID := rp.Args["userId"].(int)
	if userID != 1 {
		return
	}

	var articles []entity.Article

	article := entity.Article{
		ID:        1,
		Title:     "Golang Graphql Demo (Clean Code)",
		Subtitle:  "Build simple graphql application",
		Content:   "No Content",
		CreatedAt: time.Now().In(time.UTC),
		UserID:    int64(userID),
	}
	articles = append(articles, article)

	data = articles

	return
}
