package main

import (
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sangianpatrick/go-graphql-demo/repository"
	"github.com/sangianpatrick/go-graphql-demo/usecase"

	"github.com/sangianpatrick/go-graphql-demo/server"

	"github.com/graphql-go/graphql"
	"github.com/sangianpatrick/go-graphql-demo/handler"
	"github.com/sangianpatrick/go-graphql-demo/object"
	"github.com/sangianpatrick/go-graphql-demo/resolver"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
	db     *sql.DB
)

func main() {

	// user domain
	userRepository := repository.NewUserRepository(logger, db)
	userUsecase := usecase.NewUserUsecase(logger, userRepository)
	userResolver := resolver.NewUserResolver(userUsecase)
	userFieldConfigArgument := object.NewUserFieldConfigArgument()
	userType := object.NewUserType()

	// article domain
	articleResolver := resolver.NewArticleResolver()
	articleFieldConfigArgument := object.NewArticleFieldConfigArgument()
	articleType := object.NewArticleType()
	articleListType := graphql.NewList(articleType)

	queryType := object.NewQueryType()
	queryType.AddFieldConfig("userDetail", &graphql.Field{
		Type:        userType,
		Args:        userFieldConfigArgument,
		Resolve:     userResolver.GetUserByID,
		Description: "User detail",
	})
	queryType.AddFieldConfig("articleDetail", &graphql.Field{
		Type:        articleType,
		Args:        articleFieldConfigArgument,
		Resolve:     articleResolver.GetArticleByID,
		Description: "Article detail",
	})
	queryType.AddFieldConfig("bunchOfArticlesByUserId", &graphql.Field{
		Type:        articleListType,
		Args:        articleFieldConfigArgument,
		Resolve:     articleResolver.GetManyArticleByUserID,
		Description: "Bunch of articles",
	})

	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)

	mux := http.NewServeMux()
	handler.NewGraphqlHandler(mux, schema)

	s := server.NewServer(logrus.New(), 9999, mux)
	s.ListenAndServe()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	s.Shutdown()
}
