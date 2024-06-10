package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joaaomarques/crud-golang/src/configuration/rest_err"
	"github.com/joaaomarques/crud-golang/src/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Some fields are incorrect, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
