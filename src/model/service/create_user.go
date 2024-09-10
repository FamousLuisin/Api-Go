package service

import (
	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "createUser"))
		return nil, err
	}

	return userDomainRepository, nil
}
