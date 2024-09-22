package controller

import (
	"net/http"

	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/view"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init FindUser controller",
		zap.String("journey", "findUser"),
	)

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

	user, err := uc.Service.FindUserByIdServices(userId)

	if err != nil {
		logger.Error("Error trying to call services findUserById",
			err,
			zap.String("journey", "findUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller successfully",
		zap.String("journey", "findUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(user))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUser controller",
		zap.String("journey", "findUser"),
	)

	userEmail := c.Param("userEmail")

	user, err := uc.Service.FindUserByEmailServices(userEmail)

	if err != nil {
		logger.Error("Error trying to call services findUserByEmail",
			err,
			zap.String("journey", "findUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller successfully",
		zap.String("journey", "findUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(user))
}
