package controller

import (
	"net/http"

	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {

	userId := c.Param("userId")

	if _, err := uuid.Parse(userId); err != nil {
		logger.Error("Error trying to validate userId",
			err,
			zap.String("journey", "DeleteUser"),
		)
		errorMessage := rest_err.NewBadRequestError("UserId is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := uc.Service.DeleteUserServices(userId); err != nil {
		logger.Error("Error truing to call UpdateUser Service", err, zap.String("journey", "UpdateUser"))
		c.Status(err.Code)
		return
	}

	c.Status(http.StatusOK)
}
