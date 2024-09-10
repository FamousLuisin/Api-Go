package repository

import (
	"context"
	"database/sql"

	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/model"
)

func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{
		databaseConnection: database,
		ctx:                context.Background(),
	}
}

type userRepository struct {
	databaseConnection *sql.DB
	ctx                context.Context
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
