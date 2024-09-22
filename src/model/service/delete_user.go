package service

import "github.com/FamousLuisin/api-go/src/config/rest_err"

func (ud *userDomainService) DeleteUserServices(userId string) *rest_err.RestErr {
	return ud.userRepository.DeleteUser(userId)
}
