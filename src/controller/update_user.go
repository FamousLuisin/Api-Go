package controller

import (
	"net/http"

	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/config/validation"
	"github.com/FamousLuisin/api-go/src/controller/model/request"
	"github.com/FamousLuisin/api-go/src/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserUpdateRequest

	userId := c.Param("userId")

	if _, err := uuid.Parse(userId); err != nil {
		logger.Error("Error trying to validate userId",
			err,
			zap.String("journey", "findUser"),
		)
		errorMessage := rest_err.NewBadRequestError("UserId is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "UpdateUser"),
		)

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Age)

	err := uc.Service.UpdateUserServices(userId, domain)

	if err != nil {
		logger.Error("Error truing to call UpdateUser Service", err, zap.String("journey", "UpdateUser"))
		c.Status(err.Code)
		return
	}

	logger.Info("User update successfully",
		zap.String("userId", domain.GetId()),
		zap.String("journey", "UpdateUser"),
	)

	c.Status(http.StatusOK)
}
