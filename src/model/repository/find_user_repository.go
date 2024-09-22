package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/model"
	"go.uber.org/zap"
)

var (
	SELECT_USER_EMAIL = "SELECT id, name, age, password FROM users WHERE email = ?"
	SELECT_USER_ID    = "SELECT name, age, email, password FROM users WHERE id = ?"
	SELECT_USER_LOGIN = "SELECT id, name, age FROM users WHERE email = ? AND password = ?"
)

func (ur *userRepository) FindUserByEmail(userEmail string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository")

	stmt, err := ur.databaseConnection.PrepareContext(ur.ctx, SELECT_USER_EMAIL)

	if err != nil {
		logger.Error("Error PrepareContext", err, zap.String("journey", "findUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	var age int8
	var id, name, password string

	err = stmt.QueryRowContext(ur.ctx, userEmail).Scan(&id, &name, &age, &password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("User not found by email")
		} else {
			log.Fatal(err)
			return nil, nil
		}
	}

	fmt.Printf("User: ID=%s, Name=%s, Email=%s\n", id, name, userEmail)

	user := model.NewUserDomain(userEmail, password, name, age)
	user.SetId(id)

	return user, nil

}

func (ur *userRepository) FindUserById(userId string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository")

	stmt, err := ur.databaseConnection.PrepareContext(ur.ctx, SELECT_USER_ID)

	if err != nil {
		logger.Error("Error PrepareContext", err, zap.String("journey", "findUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	var age int8
	var email, name, password string

	err = stmt.QueryRowContext(ur.ctx, userId).Scan(&name, &age, &email, &password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("User not found by id")
		} else {
			log.Fatal(err)
			return nil, nil
		}
	}

	fmt.Printf("User: ID=%s, Name=%s, Email=%s\n", userId, name, email)

	user := model.NewUserDomain(email, password, name, age)
	user.SetId(userId)

	return user, nil

}

func (ur *userRepository) FindUserByLogin(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserByLogin repository")

	stmt, err := ur.databaseConnection.PrepareContext(ur.ctx, SELECT_USER_LOGIN)

	if err != nil {
		logger.Error("Error PrepareContext", err, zap.String("journey", "findUserByLogin"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	var age int8
	var id, name string

	err = stmt.QueryRowContext(ur.ctx, email, password).Scan(&id, &name, &age)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewForbiddenError("User or password is invalid")
		} else {
			log.Fatal(err)
			return nil, nil
		}
	}

	fmt.Printf("User: ID=%s, Name=%s, Email=%s\n", id, name, email)

	user := model.NewUserDomain(email, password, name, age)
	user.SetId(id)

	return user, nil
}
