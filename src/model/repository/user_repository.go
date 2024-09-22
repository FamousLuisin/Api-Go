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
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserById(userId string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByLogin(email, password string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(userId string) *rest_err.RestErr
}
