package repository

import (
	"context"
	"database/sql"

	"github.com/sangianpatrick/go-graphql-demo/entity"
	"github.com/sirupsen/logrus"
)

// UserRepository is a collection of behavior of user repository.
type UserRepository interface {
	FindByUserID(ctx context.Context, userID int64) (user entity.User, err error)
}

type userRepositoryImplementation struct {
	logger *logrus.Logger
	db     *sql.DB
}

// NewUserRepository is a constructor.
func NewUserRepository(logger *logrus.Logger, db *sql.DB) UserRepository {
	return &userRepositoryImplementation{
		logger: logger,
		db:     db,
	}
}

func (r *userRepositoryImplementation) FindByUserID(ctx context.Context, userID int64) (user entity.User, err error) {
	return
}
