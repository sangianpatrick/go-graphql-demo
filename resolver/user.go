package resolver

import (
	"time"

	"github.com/sangianpatrick/go-graphql-demo/usecase"

	"github.com/graphql-go/graphql"
	"github.com/sangianpatrick/go-graphql-demo/entity"
)

// UserResolver is a collection of behavior of user resolver.
type UserResolver interface {
	GetUserByID(rp graphql.ResolveParams) (data interface{}, err error)
}

type userResolverImplementation struct {
	userUsecase usecase.UserUsecase
}

// NewUserResolver is a constructor.
func NewUserResolver(userUsecase usecase.UserUsecase) UserResolver {
	return &userResolverImplementation{
		userUsecase: userUsecase,
	}
}

func (r userResolverImplementation) GetUserByID(rp graphql.ResolveParams) (data interface{}, err error) {
	id := rp.Args["id"].(int)
	if id != 1 {
		return
	}
	var user = entity.User{
		ID:        1,
		Email:     "patrick@email.com",
		Name:      "Patrick",
		CreatedAt: time.Now().In(time.UTC),
	}

	data = user

	return
}
