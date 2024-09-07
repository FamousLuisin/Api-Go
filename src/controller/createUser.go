package controller

import (
	"fmt"
	"net/http"

	"github.com/FamousLuisin/api-go/src/config/rest_err"
	"github.com/FamousLuisin/api-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect filds, error=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, userRequest)
}
