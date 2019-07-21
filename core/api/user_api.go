package api

import (
	"bytes"
	"net/http"

	"God/core/controller"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

var userController = &controller.UserController{}

func setUserApi(rootApi *gin.RouterGroup) {

	api := rootApi.Group("/user")

	api.GET("/welcome", func(ctx *gin.Context) {
		firstname := ctx.DefaultQuery("firstname", "Guest")
		lastname := ctx.Query("lastname")

		ctx.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	api.GET("/queryUserInfo", userController.GetUserInfo)
}
