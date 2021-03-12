package usecase

import (
	"context"

	"github.com/sangianpatrick/go-graphql-demo/repository"
	"github.com/sirupsen/logrus"

	"github.com/sangianpatrick/go-graphql-demo/entity"
)

// UserUsecase is an collection of behavior of user usecase.
type UserUsecase interface {
	GetByUserID(ctx context.Context, userID int64) (user entity.User, err error)
}

type userUsecaseImplementation struct {
	logger         *logrus.Logger
	userRepository repository.UserRepository
}

// NewUserUsecase is a constructor.
func NewUserUsecase(logger *logrus.Logger, userRepository repository.UserRepository) UserUsecase {
	return &userUsecaseImplementation{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (u userUsecaseImplementation) GetByUserID(ctx context.Context, userID int64) (user entity.User, err error) {
	return
}
