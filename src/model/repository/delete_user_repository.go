package repository

import (
	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"go.uber.org/zap"
)

var (
	DELETE_USER = "DELETE FROM users WHERE id = ?"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {

	logger.Info("Init DeleteUser repository")

	stmt, err := ur.databaseConnection.PrepareContext(ur.ctx, DELETE_USER)

	if err != nil {
		logger.Error("Error PrepareContext", err, zap.String("journey", "DeleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	_, err = stmt.ExecContext(ur.ctx, userId)

	if err != nil {
		logger.Error("Error ExecContext", err, zap.String("journey", "DeleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
