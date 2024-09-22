package service

import (
	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserServices(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Init UpdateUser model", zap.String("journey", "UpdateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)

	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "UpdateUser"))
		return err
	}

	return nil
}
