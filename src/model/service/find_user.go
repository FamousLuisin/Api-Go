package service

import (
	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIdServices(userId string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUser Id", zap.String("journey", "findUser"))

	return ud.userRepository.FindUserById(userId)
}

func (ud *userDomainService) FindUserByEmailServices(userEmail string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUser Email", zap.String("journey", "findUser"))

	return ud.userRepository.FindUserByEmail(userEmail)
}

func (ud *userDomainService) findUserByLogin(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUser Email", zap.String("journey", "findUser"))

	return ud.userRepository.FindUserByLogin(email, password)
}
