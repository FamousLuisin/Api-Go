package service

import (
	"fmt"

	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
