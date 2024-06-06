package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaaomarques/crud-golang/src/controller/user"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/user/:userId", user.FindUserById)
	r.GET("/user/:userEmail", user.FindUserByEmail)
	r.POST("/user/", user.CreateUser)
	r.PUT("/user/:userId", user.UpdateUser)
	r.DELETE("/user/:userId", user.DeleteUser)

}
