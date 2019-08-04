package api

import (
	"God/core/controller"

	"github.com/gin-gonic/gin"
)

var user = &controller.UserController{}

func setUserApi(rootApi *gin.RouterGroup) {

	api := rootApi.Group("/user")

	// 注册
	api.POST("/register", user.Register)
	// 密码登录
	api.POST("/login", user.LoginByPwd)
	// 验证码登录
	api.POST("/loginSms", user.LoginByMobileSms)
	//

	api.GET("/queryUserBase", user.GetUserBase)
}
