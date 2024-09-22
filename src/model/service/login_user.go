package service

import (
	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init LoginUser model", zap.String("journey", "LoginUser"))

	userDomain.EncryptPassword()
	user, err := ud.findUserByLogin(userDomain.GetEmail(), userDomain.GetPassword())

	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "LoginUser"))
		return nil, err
	}

	return user, nil
}
