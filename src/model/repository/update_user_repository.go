package repository

import (
	"fmt"
	"strings"

	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/model"
	"go.uber.org/zap"
)

var (
	UPDATE_USER = "UPDATE users SET %s WHERE id = ?"
)

func (ur *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser repository")

	var fieldsToUpdate []string
	var values []interface{}

	if userDomain.GetName() != "" {
		fieldsToUpdate = append(fieldsToUpdate, "name = ?")
		values = append(values, userDomain.GetName())
	}

	if userDomain.GetAge() != 0 {
		fieldsToUpdate = append(fieldsToUpdate, "age = ?")
		values = append(values, userDomain.GetAge())
	}

	if len(fieldsToUpdate) == 0 {
		return rest_err.NewBadRequestError("No fields to update")
	}

	query := fmt.Sprintf(UPDATE_USER, strings.Join(fieldsToUpdate, ", "))

	values = append(values, userId)

	stmt, err := ur.databaseConnection.PrepareContext(ur.ctx, query)

	if err != nil {
		logger.Error("Error PrepareContext", err, zap.String("journey", "UpdateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ur.ctx, values...)

	if err != nil {
		logger.Error("Error ExecContext", err, zap.String("journey", "UpdateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
