package controller

import (
	"net/http"

	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/validation"
	"github.com/FamousLuisin/api-go/src/controller/model/request"
	"github.com/FamousLuisin/api-go/src/model"
	"github.com/FamousLuisin/api-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	domainResult, err := uc.Service.CreateUser(domain)

	if err != nil {
		logger.Error("Error truing to call CreateUser Service", err, zap.String("journey", "createUser"))
		return
	}

	logger.Info("User created successfully",
		zap.String("userId", domain.GetId()),
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
