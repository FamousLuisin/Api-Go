package service

import (
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/model"
	"github.com/FamousLuisin/api-go/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{
		userRepository: userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserServices(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByEmailServices(userEmail string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIdServices(userId string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserServices(string) *rest_err.RestErr
	LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)

	findUserByLogin(email, password string) (model.UserDomainInterface, *rest_err.RestErr)
}
