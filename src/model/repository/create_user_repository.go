package repository

import (
	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	INSERT_USER = "INSERT INTO users (id, name, email, password, age) VALUES (?, ?, ?, ?, ?)"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser repository")

	stmt, err := ur.databaseConnection.PrepareContext(ur.ctx, INSERT_USER)

	if err != nil {
		logger.Error("Error PrepareContext", err, zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	id := uuid.New()
	name := userDomain.GetName()
	email := userDomain.GetEmail()
	password := userDomain.GetPassword()
	age := userDomain.GetAge()

	_, err = stmt.ExecContext(ur.ctx, id.String(), name, email, password, age)

	if err != nil {
		logger.Error("Error ExecContext", err, zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetId(id.String())

	return userDomain, nil
}
