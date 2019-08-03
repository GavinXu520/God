package api

import (
	"God/core/controller"

	"github.com/gin-gonic/gin"
)

var userController = &controller.UserController{}

func setUserApi(rootApi *gin.RouterGroup) {

	api := rootApi.Group("/user")

	// 注册
	api.POST("/register", userController.Register)
	// 密码登录
	api.POST("/login", userController.LoginByPwd)

	api.GET("/queryUserBase", userController.GetUserBase)
}
