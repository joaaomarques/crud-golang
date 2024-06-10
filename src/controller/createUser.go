package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/joaaomarques/crud-golang/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("Wrong route called!")
	c.JSON(err.Code, err)
}
