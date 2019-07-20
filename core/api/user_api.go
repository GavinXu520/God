package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setUserApi(rootApi *gin.RouterGroup) {

	api := rootApi.Group("/user")

	api.GET("/welcome", func(ctx *gin.Context) {
		firstname := ctx.DefaultQuery("firstname", "Guest")
		lastname := ctx.Query("lastname")

		ctx.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

}
