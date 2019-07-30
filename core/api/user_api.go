package api

import (
	"God/core/controller"

	"github.com/gin-gonic/gin"
)

var userController = &controller.UserController{}

func setUserApi(rootApi *gin.RouterGroup) {

	api := rootApi.Group("/user")

	api.POST("/register", userController.GetUserInfo)
}
