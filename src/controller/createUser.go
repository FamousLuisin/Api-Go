package controller

import (
	"log"
	"net/http"

	"github.com/FamousLuisin/api-go/src/config/validation"
	"github.com/FamousLuisin/api-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=\n%s", err.Error())

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, userRequest)
}
